package models

import "gorm.io/gorm"

type RoleMenu struct {
	gorm.Model
	RoleId uint `gorm:"column:role_id;type:int(11);" json:"role_id"` // 角色ID
	MenuId uint `gorm:"column:menu_id;type:int(11);" json:"menu_id"` // 菜单ID
}

func (table *RoleMenu) TableName() string {
	return "role_menu"
}
