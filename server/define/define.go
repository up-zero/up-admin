package define

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	Port           = ":9090"
	FrameName      = "UpAdmin"
	DateTimeLayout = "2006-01-02 15:04:05"
	// DefaultSize 默认每页查询20条数据
	DefaultSize = 20
	// UpAdminDSN 数据库连接信息配置
	UpAdminDSN = ""
	// JwtKey 密钥（建议修改）
	JwtKey = "up-admin"
	// TokenExpire token 有效期，7天
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// RefreshTokenExpire 刷新 token 有效期，14天
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
)

// Redis 相关配置
var (
	RedisAddr     = "192.168.1.8:6379"
	RedisUsername = ""
	RedisPassword = ""
	// RedisRoleAdminPrefix 判断角色是否是超管的前缀
	RedisRoleAdminPrefix = "ADMIN-"
	// RedisMenuPrefix 菜单
	RedisMenuPrefix = "MENU"
	// RedisFuncPrefix 功能的前缀
	RedisFuncPrefix = "FUNC-"
)

func InitEnv() {
	UpAdminDSN = os.Getenv("UpAdminDSN")
}

type UserClaim struct {
	Id           uint
	Identity     string
	Name         string
	RoleIdentity string // 角色唯一标识
	IsAdmin      bool   // 是否是超管
	jwt.StandardClaims
}
