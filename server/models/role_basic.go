package models

import "gorm.io/gorm"

type RoleBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`
	IsAdmin  int8   `gorm:"column:is_admin;type:tinyint(1);default:0;" json:"is_admin"` // 是否是超管【0-否 1-是】
}

func (table *RoleBasic) TableName() string {
	return "role_basic"
}
