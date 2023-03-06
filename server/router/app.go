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
	// 获取功能列表
	loginAuth.GET("/functions", service.Functions)
	// 获取用户信息
	loginAuth.GET("/user/info", service.UserInfo)
	// 修改密码
	loginAuth.PUT("/user/password/change", service.UserPasswordChange)

	// 会鉴权的接口
	auth := loginAuth.Use(middleware.FuncAuthCheck())

	// ---------------- BEGIN - 设置 ----------------
	// 角色管理
	// 角色列表
	auth.GET("/set/role/list", service.SetRoleList)
	// 角色详情
	auth.GET("/set/role/detail", service.SetRoleDetail)
	// 修改角色的管理员身份
	auth.PUT("/set/role/update/admin", service.SetRoleUpdateAdmin)
	// 新增角色
	auth.POST("/set/role/create", service.SetRoleCreate)
	// 删除角色
	auth.DELETE("/set/role/delete", service.SetRoleDelete)
	// 修改角色
	auth.PUT("/set/role/update", service.SetRoleUpdate)

	// 菜单功能管理
	// 菜单列表
	auth.GET("/set/menu/list", service.SetMenuList)
	// 功能列表
	auth.GET("/set/func/list", service.SetFuncList)
	auth.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Success",
		})
	})

	// 管理员管理
	// 管理员列表
	auth.GET("/set/user/list", service.SetUserList)
	// ---------------- END - 设置 ----------------

	// ---------------- BEGIN - dev ----------------
	// 新增菜单
	auth.POST("/dev/menu/add", service.DevMenuAdd)
	// 修改菜单
	auth.PUT("/dev/menu/update", service.DevMenuUpdate)
	// 删除菜单
	auth.DELETE("/dev/menu/delete", service.DevMenuDelete)
	// 新增功能
	auth.POST("/dev/func/add", service.DevFuncAdd)
	// 修改功能
	auth.PUT("/dev/func/update", service.DevFuncUpdate)
	// 删除功能
	auth.DELETE("/dev/func/delete", service.DevFuncDelete)
	// ---------------- END - dev ----------------

	return r
}
