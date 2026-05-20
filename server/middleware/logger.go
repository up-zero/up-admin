package middleware

import (
	"errors"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"gitee.com/up-zero/up-admin/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func statusEvent(status int) *zerolog.Event {
	current := logger.Current()
	switch {
	case status >= http.StatusInternalServerError:
		return current.Error()
	case status >= http.StatusBadRequest:
		return current.Warn()
	default:
		return current.Info()
	}
}

func requestEvent(c *gin.Context, status int, latencyMs int64) *zerolog.Event {
	event := statusEvent(status).
		Str("type", "access").
		Int("status", status).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("route", c.FullPath()).
		Str("query", c.Request.URL.RawQuery).
		Str("client_ip", c.ClientIP()).
		Str("user_agent", c.Request.UserAgent()).
		Int64("latency_ms", latencyMs).
		Int("body_size", c.Writer.Size())
	if requestID := strings.TrimSpace(c.GetHeader("X-Request-Id")); requestID != "" {
		event = event.Str("request_id", requestID)
	}
	if len(c.Errors) > 0 {
		event = event.Str("errors", c.Errors.String())
	}
	return event
}

func isBrokenPipeError(err error) bool {
	var netErr *net.OpError
	if !errors.As(err, &netErr) {
		return false
	}
	if netErr.Err == nil {
		return false
	}
	message := strings.ToLower(netErr.Err.Error())
	return strings.Contains(message, "broken pipe") || strings.Contains(message, "connection reset by peer")
}

// GinLogger 记录 Gin 请求访问日志。
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latencyMs := time.Since(start).Milliseconds()
		status := c.Writer.Status()
		requestEvent(c, status, latencyMs).Msg("http request")
	}
}

// GinRecovery 捕获 Gin 请求中的 panic 并记录结构化日志。
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				err, _ := recovered.(error)
				if err != nil && isBrokenPipeError(err) {
					c.Status(499)
					requestEvent(c, 499, 0).
						Str("type", "panic").
						Interface("panic", recovered).
						Msg("request terminated by client connection")
					c.Abort()
					return
				}

				event := statusEvent(http.StatusInternalServerError).
					Str("type", "panic").
					Str("method", c.Request.Method).
					Str("path", c.Request.URL.Path).
					Str("route", c.FullPath()).
					Str("query", c.Request.URL.RawQuery).
					Str("client_ip", c.ClientIP()).
					Str("user_agent", c.Request.UserAgent()).
					Str("stack", string(debug.Stack())).
					Interface("panic", recovered)
				if requestID := strings.TrimSpace(c.GetHeader("X-Request-Id")); requestID != "" {
					event = event.Str("request_id", requestID)
				}
				event.Msg("panic recovered")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
