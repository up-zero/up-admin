package models

import (
	"gitee.com/up-zero/up-admin/define"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGormDB() {
	db, err := gorm.Open(mysql.Open(define.UpAdminDSN))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&UserBasic{})
	if err != nil {
		panic(err)
	}
	DB = db
}
