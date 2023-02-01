package models

import "gorm.io/gorm"

type RoleBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`
	IsAdmin  int8   `gorm:"column:is_admin;type:tinyint(1);default:0;" json:"is_admin"` // 是否是超管【0-否 1-是】
	Sort     int64  `gorm:"column:sort;type:int(11);default:0;" json:"sort"`            // 排序，序号越少越靠前
}

func (table *RoleBasic) TableName() string {
	return "role_basic"
}

// GetRoleList 获取角色列表
func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(new(RoleBasic)).Select("identity, name, is_admin, sort, created_at, updated_at")
	if keyword != "" {
		tx.Where("name LIKE ?", "%"+keyword+"%")
	}
	tx.Order("sort ASC")
	return tx
}

// GetRoleBasic 获取角色详情
func GetRoleBasic(identity string) (*RoleBasic, error) {
	rb := new(RoleBasic)
	err := DB.Model(new(RoleBasic)).Where("identity = ?", identity).First(rb).Error
	return rb, err
}
