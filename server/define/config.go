package define

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	FrameName      string `mapstructure:"frame_name"`
	DateTimeLayout string `mapstructure:"date_time_layout"`
	DefaultSize    int    `mapstructure:"default_size"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type LogConfig struct {
	Path       string `mapstructure:"path"`        // 日志文件路径
	Level      string `mapstructure:"level"`       // 日志级别（debug、info、warn、error）
	MaxSize    int    `mapstructure:"max_size"`    // 单个日志文件最大体积，单位 MB
	MaxBackups int    `mapstructure:"max_backups"` // 历史日志文件保留数量
	MaxAge     int    `mapstructure:"max_age"`     // 历史日志保留天数
	Compress   bool   `mapstructure:"compress"`    // 是否压缩旧日志文件
	Console    bool   `mapstructure:"console"`     // 是否输出到控制台
}

type RedisConfig struct {
	Addr            string `mapstructure:"addr"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	RoleAdminPrefix string `mapstructure:"role_admin_prefix"`
	MenuPrefix      string `mapstructure:"menu_prefix"`
	FuncPrefix      string `mapstructure:"func_prefix"`
}

type JWTConfig struct {
	Key             string `mapstructure:"key"`
	TokenTTL        string `mapstructure:"token_ttl"`
	RefreshTokenTTL string `mapstructure:"refresh_token_ttl"`
}

type ReloadHandler func(oldConfig, newConfig Config) error

type runtimeConfig struct {
	settings        Config
	tokenTTL        time.Duration
	refreshTokenTTL time.Duration
}

var (
	configMu       sync.RWMutex
	appSettings    Config
	configFilePath string
	configViper    *viper.Viper
	reloadHandlers []ReloadHandler
)

func InitConfig(path string) error {
	configFilePath = resolveConfigFilePath(path)
	v := newViper(configFilePath)

	initialized, err := ensureDefaultConfigFile(v, configFilePath)
	if err != nil {
		return err
	}

	rc, err := loadRuntimeConfig(v, configFilePath)
	if err != nil {
		if initialized {
			return fmt.Errorf("已初始化默认配置文件 %q，请先完善配置: %w", configFilePath, err)
		}
		return err
	}

	configMu.Lock()
	configViper = v
	appSettings = rc.settings
	configMu.Unlock()
	applyConfig(rc)

	if initialized {
		log.Printf("初始化默认配置成功: %s", configFilePath)
	}
	return nil
}

func CurrentConfig() Config {
	configMu.RLock()
	defer configMu.RUnlock()
	return appSettings
}

func RegisterConfigReloadHandler(handler ReloadHandler) {
	if handler == nil {
		return
	}
	configMu.Lock()
	reloadHandlers = append(reloadHandlers, handler)
	configMu.Unlock()
}

func WatchConfig() error {
	configMu.RLock()
	v := configViper
	path := configFilePath
	configMu.RUnlock()
	if v == nil {
		return fmt.Errorf("配置尚未初始化")
	}

	v.OnConfigChange(func(event fsnotify.Event) {
		if err := reloadConfig(v, path); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log.Printf("配置文件暂不可用，等待下次变更: %v", err)
				return
			}
			log.Printf("配置热重载失败: %v", err)
			return
		}
		log.Printf("配置热重载成功: %s (%s)", event.Name, event.Op.String())
	})
	v.WatchConfig()
	log.Printf("已开启配置热重载: %s", path)
	return nil
}

func newViper(configFilePath string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(configFilePath)
	v.SetConfigType("yaml")
	setDefaults(v)
	return v
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.frame_name", "UpAdmin")
	v.SetDefault("app.date_time_layout", "2006-01-02 15:04:05")
	v.SetDefault("app.default_size", 20)
	v.SetDefault("server.port", ":9090")
	v.SetDefault("database.dsn", "")
	v.SetDefault("log.path", "logs/up-admin.log")
	v.SetDefault("log.level", "info")
	v.SetDefault("log.max_size", 100)
	v.SetDefault("log.max_backups", 7)
	v.SetDefault("log.max_age", 30)
	v.SetDefault("log.compress", false)
	v.SetDefault("log.console", true)
	v.SetDefault("redis.addr", "127.0.0.1:6379")
	v.SetDefault("redis.username", "")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.role_admin_prefix", "ADMIN-")
	v.SetDefault("redis.menu_prefix", "MENU")
	v.SetDefault("redis.func_prefix", "FUNC-")
	v.SetDefault("jwt.key", "up-admin")
	v.SetDefault("jwt.token_ttl", "168h")
	v.SetDefault("jwt.refresh_token_ttl", "336h")
}

func ensureDefaultConfigFile(v *viper.Viper, configFilePath string) (bool, error) {
	if _, err := os.Stat(configFilePath); err == nil {
		return false, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("检查配置文件失败 %q: %w", configFilePath, err)
	}

	configDir := filepath.Dir(configFilePath)
	if configDir != "." && configDir != "" {
		if err := os.MkdirAll(configDir, 0o755); err != nil {
			return false, fmt.Errorf("创建配置目录失败 %q: %w", configDir, err)
		}
	}
	if err := v.SafeWriteConfigAs(configFilePath); err != nil {
		return false, fmt.Errorf("初始化默认配置文件失败 %q: %w", configFilePath, err)
	}
	return true, nil
}

func resolveConfigFilePath(path string) string {
	configFilePath := strings.TrimSpace(path)
	if configFilePath == "" {
		configFilePath = defaultConfigFile
	}
	return configFilePath
}

func loadRuntimeConfig(v *viper.Viper, configFilePath string) (runtimeConfig, error) {
	if err := v.ReadInConfig(); err != nil {
		return runtimeConfig{}, fmt.Errorf("读取配置文件失败 %q: %w", configFilePath, err)
	}

	cfg := Config{}
	if err := v.Unmarshal(&cfg); err != nil {
		return runtimeConfig{}, fmt.Errorf("解析配置文件失败 %q: %w", configFilePath, err)
	}

	return normalizeConfig(cfg)
}

func reloadConfig(v *viper.Viper, configFilePath string) error {
	rc, err := loadRuntimeConfig(v, configFilePath)
	if err != nil {
		return err
	}

	configMu.RLock()
	oldConfig := appSettings
	handlers := append([]ReloadHandler(nil), reloadHandlers...)
	configMu.RUnlock()

	for _, handler := range handlers {
		if err := handler(oldConfig, rc.settings); err != nil {
			return err
		}
	}

	configMu.Lock()
	appSettings = rc.settings
	configMu.Unlock()
	applyConfig(rc)
	return nil
}

func normalizeConfig(cfg Config) (runtimeConfig, error) {
	port := strings.TrimSpace(cfg.Server.Port)
	if port == "" {
		return runtimeConfig{}, fmt.Errorf("server.port 不能为空")
	}
	if !strings.Contains(port, ":") {
		port = ":" + port
	}

	dsn := strings.TrimSpace(cfg.Database.DSN)
	if dsn == "" {
		return runtimeConfig{}, fmt.Errorf("database.dsn 不能为空")
	}

	logPath := strings.TrimSpace(cfg.Log.Path)
	if logPath == "" {
		logPath = "logs/up-admin.log"
	}
	logLevel := strings.TrimSpace(cfg.Log.Level)
	if logLevel == "" {
		logLevel = "info"
	}
	logMaxSize := cfg.Log.MaxSize
	if logMaxSize <= 0 {
		logMaxSize = 100
	}
	logMaxBackups := cfg.Log.MaxBackups
	if logMaxBackups <= 0 {
		logMaxBackups = 7
	}
	logMaxAge := cfg.Log.MaxAge
	if logMaxAge <= 0 {
		logMaxAge = 30
	}

	jwtKey := strings.TrimSpace(cfg.JWT.Key)
	if jwtKey == "" {
		return runtimeConfig{}, fmt.Errorf("jwt.key 不能为空")
	}

	tokenTTL, err := time.ParseDuration(strings.TrimSpace(cfg.JWT.TokenTTL))
	if err != nil {
		return runtimeConfig{}, fmt.Errorf("解析 jwt.token_ttl 失败: %w", err)
	}
	refreshTokenTTL, err := time.ParseDuration(strings.TrimSpace(cfg.JWT.RefreshTokenTTL))
	if err != nil {
		return runtimeConfig{}, fmt.Errorf("解析 jwt.refresh_token_ttl 失败: %w", err)
	}
	if tokenTTL <= 0 {
		return runtimeConfig{}, fmt.Errorf("jwt.token_ttl 必须大于 0")
	}
	if refreshTokenTTL <= 0 {
		return runtimeConfig{}, fmt.Errorf("jwt.refresh_token_ttl 必须大于 0")
	}

	frameName := strings.TrimSpace(cfg.App.FrameName)
	if frameName == "" {
		frameName = "UpAdmin"
	}
	dateTimeLayout := strings.TrimSpace(cfg.App.DateTimeLayout)
	if dateTimeLayout == "" {
		dateTimeLayout = "2006-01-02 15:04:05"
	}
	defaultSize := cfg.App.DefaultSize
	if defaultSize <= 0 {
		defaultSize = 20
	}
	redisAddr := strings.TrimSpace(cfg.Redis.Addr)
	if redisAddr == "" {
		redisAddr = "127.0.0.1:6379"
	}
	redisRoleAdminPrefix := strings.TrimSpace(cfg.Redis.RoleAdminPrefix)
	if redisRoleAdminPrefix == "" {
		redisRoleAdminPrefix = "ADMIN-"
	}
	redisMenuPrefix := strings.TrimSpace(cfg.Redis.MenuPrefix)
	if redisMenuPrefix == "" {
		redisMenuPrefix = "MENU"
	}
	redisFuncPrefix := strings.TrimSpace(cfg.Redis.FuncPrefix)
	if redisFuncPrefix == "" {
		redisFuncPrefix = "FUNC-"
	}

	return runtimeConfig{
		settings: Config{
			App: AppConfig{
				FrameName:      frameName,
				DateTimeLayout: dateTimeLayout,
				DefaultSize:    defaultSize,
			},
			Server: ServerConfig{
				Port: port,
			},
			Database: DatabaseConfig{
				DSN: dsn,
			},
			Log: LogConfig{
				Path:       logPath,
				Level:      logLevel,
				MaxSize:    logMaxSize,
				MaxBackups: logMaxBackups,
				MaxAge:     logMaxAge,
				Compress:   cfg.Log.Compress,
				Console:    cfg.Log.Console,
			},
			Redis: RedisConfig{
				Addr:            redisAddr,
				Username:        strings.TrimSpace(cfg.Redis.Username),
				Password:        strings.TrimSpace(cfg.Redis.Password),
				RoleAdminPrefix: redisRoleAdminPrefix,
				MenuPrefix:      redisMenuPrefix,
				FuncPrefix:      redisFuncPrefix,
			},
			JWT: JWTConfig{
				Key:             jwtKey,
				TokenTTL:        strings.TrimSpace(cfg.JWT.TokenTTL),
				RefreshTokenTTL: strings.TrimSpace(cfg.JWT.RefreshTokenTTL),
			},
		},
		tokenTTL:        tokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}, nil
}

func applyConfig(rc runtimeConfig) {
	Port = rc.settings.Server.Port
	FrameName = rc.settings.App.FrameName
	DateTimeLayout = rc.settings.App.DateTimeLayout
	DefaultSize = rc.settings.App.DefaultSize
	JwtKey = rc.settings.JWT.Key
	TokenTTL = rc.tokenTTL
	RefreshTokenTTL = rc.refreshTokenTTL
	RedisRoleAdminPrefix = rc.settings.Redis.RoleAdminPrefix
	RedisMenuPrefix = rc.settings.Redis.MenuPrefix
	RedisFuncPrefix = rc.settings.Redis.FuncPrefix
}
