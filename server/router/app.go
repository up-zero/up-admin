package router

import (
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/service/login"
	"github.com/gin-gonic/gin"
)

func App(r *gin.Engine) {
	// 登录
	// 用户名、密码登录
	r.POST("/login/password", login.Password)
	r.Run(define.Port)
}
