package models

import "gorm.io/gorm"

type RoleFunction struct {
	gorm.Model
	RoleId     uint `gorm:"column:role_id;type:int(11);" json:"role_id"`         // 角色ID
	FunctionId uint `gorm:"column:function_id;type:int(11);" json:"function_id"` // 功能ID
}

func (table *RoleFunction) TableName() string {
	return "role_function"
}
