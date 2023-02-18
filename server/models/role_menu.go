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

// GetRoleMenus 获取指定角色的菜单列表
func GetRoleMenus(roleIdentity string, isAdmin bool) (*gorm.DB, error) {
	tx := new(gorm.DB)
	if isAdmin {
		tx = DB.Model(new(MenuBasic)).Select("id, parent_id, identity, name, web_icon, sort, path").Order("sort ASC")
	} else {
		roleBasic := new(RoleBasic)
		err := DB.Model(new(RoleBasic)).Select("id").Where("identity = ?", roleIdentity).Find(roleBasic).Error
		if err != nil {
			return nil, err
		}
		tx = DB.Model(new(RoleMenu)).Select("mb.id, mb.parent_id, mb.identity, mb.name, mb.web_icon, mb.sort, mb.path").
			Joins("LEFT JOIN menu_basic mb ON mb.id = role_menu.menu_id").
			Where("role_menu.role_id = ?", roleBasic.ID).Order("mb.sort ASC")
	}
	return tx, nil
}

// GetRoleMenuIdentity 获取指定角色的菜单唯一标识
func GetRoleMenuIdentity(roleId uint, isAdmin bool) ([]string, error) {
	tx := new(gorm.DB)
	data := make([]string, 0)
	if isAdmin {
		tx = DB.Model(new(MenuBasic)).Select("identity").Order("sort ASC")
	} else {
		tx = DB.Model(new(RoleMenu)).Select("mb.identity").
			Joins("LEFT JOIN menu_basic mb ON mb.id = role_menu.menu_id").
			Where("role_menu.role_id = ?", roleId).Order("mb.sort ASC")
	}
	err := tx.Scan(&data).Error
	return data, err
}
