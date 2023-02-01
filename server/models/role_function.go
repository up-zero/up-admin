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

// GetAuthFuncUri 获取给特定角色授权功能的URI
func GetAuthFuncUri(roleIdentity string) (map[string]interface{}, error) {
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

// GetRoleFunctionIdentity 获取指定角色的功能唯一标识
func GetRoleFunctionIdentity(roleId uint, isAdmin bool) ([]string, error) {
	tx := new(gorm.DB)
	data := make([]string, 0)
	if isAdmin {
		tx = DB.Model(new(FunctionBasic)).Select("identity").Order("sort ASC")
	} else {
		tx = DB.Model(new(RoleFunction)).Select("fb.identity").
			Joins("LEFT JOIN function_basic fb ON fb.id = role_function.function_id").
			Where("role_function.role_id = ?", roleId).Order("fb.sort ASC")
	}
	err := tx.Scan(&data).Error
	return data, err
}
