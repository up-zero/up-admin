package models

import "gorm.io/gorm"

type RoleBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`
}

func (table *RoleBasic) TableName() string {
	return "role_basic"
}
