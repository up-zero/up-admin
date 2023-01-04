package router

import (
	"gitee.com/up-zero/up-admin/middleware"
	"gitee.com/up-zero/up-admin/service"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// 登录
	// 用户名、密码登录
	r.POST("/login/password", service.Password)

	// 登录信息认证
	loginAuth := r.Group("/").Use(middleware.LoginAuthCheck())

	// 获取菜单列表
	loginAuth.GET("/menus", service.Menus)

	return r
}
