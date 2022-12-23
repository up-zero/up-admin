package router

import (
	"gitee.com/up-zero/up-admin/middleware"
	"gitee.com/up-zero/up-admin/service/login"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// 登录
	// 用户名、密码登录
	r.POST("/login/password", login.Password)
	return r
}
