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
	// UpAdminDSN 数据库连接信息配置
	UpAdminDSN = ""
	// JwtKey 密钥
	JwtKey = "up-admin"
	// TokenExpire token 有效期，7天
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// RefreshTokenExpire 刷新 token 有效期，14天
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
)

func InitEnv() {
	UpAdminDSN = os.Getenv("UpAdminDSN")
}

type UserClaim struct {
	Id       uint
	Identity string
	Name     string
	jwt.StandardClaims
}