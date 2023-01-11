package router

import (
	"gitee.com/up-zero/up-admin/middleware"
	"gitee.com/up-zero/up-admin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	// 登录
	// 用户名、密码登录
	r.POST("/login/password", service.LoginPassword)

	// 登录信息认证
	loginAuth := r.Group("/").Use(middleware.LoginAuthCheck())

	// 获取菜单列表
	loginAuth.GET("/menus", service.Menus)
	// 获取用户信息
	loginAuth.GET("/user/info", service.UserInfo)
	// 修改密码
	loginAuth.PUT("/user/password/change", service.UserPasswordChange)

	// 会鉴权的接口
	auth := loginAuth.Use(middleware.FuncAuthCheck())

	// 设置
	// 角色列表
	auth.GET("/set/role/list", service.SetRoleList)
	// 修改角色的管理员身份
	auth.PUT("/set/role/update/admin", service.SetRoleUpdateAdmin)
	// TODO：角色管理、菜单功能管理
	auth.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Success",
		})
	})

	return r
}
