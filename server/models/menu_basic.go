package models

import "gorm.io/gorm"

type MenuBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ParentId uint   `gorm:"column:parent_id;type:int(11);" json:"parent_id"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
}

func (table *MenuBasic) TableName() string {
	return "menu_basic"
}
