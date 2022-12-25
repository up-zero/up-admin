package models

import (
	"gitee.com/up-zero/up-admin/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
