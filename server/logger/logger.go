package logger

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gitee.com/up-zero/up-admin/define"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	loggerMu       sync.RWMutex
	appLogger      zerolog.Logger
	rotatingWriter *lumberjack.Logger
)

func init() {
	appLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}).With().Timestamp().Str("app", "UpAdmin").Logger()
	redirectStdLogger()
}

func InitLogger(cfg define.Config) error {
	logLevel := strings.TrimSpace(strings.ToLower(cfg.Log.Level))
	if logLevel == "" {
		logLevel = "info"
	}
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("解析 log.level 失败: %w", err)
	}

	timeLayout := strings.TrimSpace(cfg.App.DateTimeLayout)
	if timeLayout == "" {
		timeLayout = "2006-01-02 15:04:05"
	}
	frameName := strings.TrimSpace(cfg.App.FrameName)
	if frameName == "" {
		frameName = "UpAdmin"
	}

	writer, fileWriter, err := buildLogWriter(cfg.Log, timeLayout)
	if err != nil {
		return err
	}

	logger := zerolog.New(writer).Level(level).With().Timestamp().Str("app", frameName).Logger()

	loggerMu.Lock()
	prevWriter := rotatingWriter
	appLogger = logger
	rotatingWriter = fileWriter
	loggerMu.Unlock()

	redirectStdLogger()

	if prevWriter != nil && prevWriter != fileWriter {
		_ = prevWriter.Close()
	}
	return nil
}

func buildLogWriter(cfg define.LogConfig, timeLayout string) (io.Writer, *lumberjack.Logger, error) {
	outputs := make([]io.Writer, 0, 2)
	logPath := strings.TrimSpace(cfg.Path)
	if cfg.Console || logPath == "" {
		outputs = append(outputs, zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeLayout})
	}

	var fileWriter *lumberjack.Logger
	if logPath != "" {
		cleanPath := filepath.Clean(logPath)
		logDir := filepath.Dir(cleanPath)
		if logDir != "." && logDir != "" {
			if err := os.MkdirAll(logDir, 0o755); err != nil {
				return nil, nil, fmt.Errorf("创建日志目录失败 %q: %w", logDir, err)
			}
		}
		fileWriter = &lumberjack.Logger{
			Filename:   cleanPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
			LocalTime:  true,
		}
		outputs = append(outputs, fileWriter)
	}

	if len(outputs) == 0 {
		outputs = append(outputs, zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeLayout})
	}
	if len(outputs) == 1 {
		return outputs[0], fileWriter, nil
	}
	return zerolog.MultiLevelWriter(outputs...), fileWriter, nil
}

func redirectStdLogger() {
	stdlog.SetFlags(0)
	stdlog.SetOutput(stdLogWriter{})
}

type stdLogWriter struct{}

func (stdLogWriter) Write(p []byte) (int, error) {
	message := strings.TrimSpace(string(p))
	if message == "" {
		return len(p), nil
	}
	logger := currentLogger()
	logger.Info().Msg(message)
	return len(p), nil
}

func currentLogger() zerolog.Logger {
	loggerMu.RLock()
	defer loggerMu.RUnlock()
	return appLogger
}

// Current 返回当前日志实例，供其他包复用统一日志配置。
func Current() zerolog.Logger {
	return currentLogger()
}

func formatMessage(format string, v ...interface{}) string {
	if len(v) == 0 {
		return format
	}
	return fmt.Sprintf(format, v...)
}

func caller(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", filepath.ToSlash(file), line)
}

// Info INFO
func Info(format string, v ...interface{}) {
	logger := currentLogger()
	logger.Info().Msg(formatMessage(format, v...))
}

// Error ERROR
func Error(format string, v ...interface{}) {
	logger := currentLogger()
	event := logger.Error()
	if location := caller(2); location != "" {
		event = event.Str("caller", location)
	}
	event.Msg(formatMessage(format, v...))
}
