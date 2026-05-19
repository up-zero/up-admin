package main

import (
	"flag"
	"log"

	"gitee.com/up-zero/up-admin/define"
	"gitee.com/up-zero/up-admin/models"
	"gitee.com/up-zero/up-admin/router"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "配置文件路径")
	flag.Parse()

	// 初始化配置
	if err := define.InitConfig(configFile); err != nil {
		panic(err)
	}
	// 初始化数据源
	if err := models.ReloadDatastores(define.CurrentConfig()); err != nil {
		panic(err)
	}
	define.RegisterConfigReloadHandler(func(oldConfig, newConfig define.Config) error {
		if oldConfig.Database == newConfig.Database && oldConfig.Redis == newConfig.Redis {
			if oldConfig.Server.Port != newConfig.Server.Port {
				log.Printf("检测到 server.port 变更为 %s，将在下次重启后生效", newConfig.Server.Port)
			}
			return nil
		}
		if err := models.ReloadDatastores(newConfig); err != nil {
			return err
		}
		if oldConfig.Server.Port != newConfig.Server.Port {
			log.Printf("检测到 server.port 变更为 %s，将在下次重启后生效", newConfig.Server.Port)
		}
		return nil
	})
	if err := define.WatchConfig(); err != nil {
		panic(err)
	}

	r := router.App()
	if err := r.Run(define.Port); err != nil {
		panic(err)
	}
}
