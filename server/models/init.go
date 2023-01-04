package models

import (
	"gitee.com/up-zero/up-admin/define"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RDB *redis.Client

func NewGormDB() {
	db, err := gorm.Open(mysql.Open(define.UpAdminDSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&UserBasic{}, &RoleBasic{}, &RoleMenu{}, &RoleFunction{}, &MenuBasic{}, &FunctionBasic{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func NewRedisDB() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     define.RedisAddr,
		Username: define.RedisUsername,
		Password: define.RedisPassword,
	})
}
