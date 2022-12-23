package main

import (
	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/models"
	"gitee.com/up-zero/up-admin/router"
)

func main() {
	// 初始化环境变量
	define.InitEnv()
	// 初始化 gorm.DB
	models.NewGormDB()

	r := router.App()
	r.Run(define.Port)
}
