package main

import (
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/models"
	"gitee.com/up-zero/up-admin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化环境变量
	define.InitEnv()
	// 初始化 gorm.DB
	models.NewGormDB()
	r := gin.Default()
	router.App(r)
}
