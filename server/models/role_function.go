package models

import (
	"gorm.io/gorm"
)

type RoleFunction struct {
	gorm.Model
	RoleId     uint `gorm:"column:role_id;type:int(11);" json:"role_id"`         // 角色ID
	FunctionId uint `gorm:"column:function_id;type:int(11);" json:"function_id"` // 功能ID
}

func (table *RoleFunction) TableName() string {
	return "role_function"
}

type authFuncReply struct {
	Uri string `json:"uri"`
}

// GetAuthFunc 获取给特定角色授权的功能
func GetAuthFunc(roleIdentity string) (map[string]interface{}, error) {
	roleBasic := new(RoleBasic)
	err := DB.Model(new(RoleBasic)).Select("id").Where("identity = ?", roleIdentity).Find(roleBasic).Error
	if err != nil {
		return nil, err
	}
	afr := make([]*authFuncReply, 0)
	err = DB.Model(new(RoleFunction)).Select("fb.uri").
		Joins("LEFT JOIN function_basic fb ON fb.id = role_function.function_id").
		Where("role_function.role_id = ?", roleBasic.ID).Find(&afr).Error
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	for _, v := range afr {
		data[v.Uri] = "1"
	}
	return data, nil
}
