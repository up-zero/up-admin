package define

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	Port           = ":9090"
	FrameName      = "UpAdmin"
	DateTimeLayout = "2006-01-02 15:04:05"
	// DefaultSize 默认每页查询20条数据
	DefaultSize = 20
	// JwtKey 密钥（建议修改）
	JwtKey = "up-admin"
	// TokenTTL token 有效期，7天
	TokenTTL = time.Hour * 24 * 7
	// RefreshTokenTTL 刷新 token 有效期，14天
	RefreshTokenTTL = time.Hour * 24 * 14
)

// Redis 相关配置
var (
	// RedisRoleAdminPrefix 判断角色是否是超管的前缀
	RedisRoleAdminPrefix = "ADMIN-"
	// RedisMenuPrefix 菜单
	RedisMenuPrefix = "MENU"
	// RedisFuncPrefix 功能的前缀
	RedisFuncPrefix = "FUNC-"
)

func TokenExpireAt() int64 {
	return time.Now().Add(TokenTTL).Unix()
}

func RefreshTokenExpireAt() int64 {
	return time.Now().Add(RefreshTokenTTL).Unix()
}

type UserClaim struct {
	Id           uint
	Identity     string
	Name         string
	RoleIdentity string // 角色唯一标识
	IsAdmin      bool   // 是否是超管
	jwt.StandardClaims
}
