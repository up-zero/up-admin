package models

import "gorm.io/gorm"

type FunctionBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	MenuId   uint   `gorm:"column:menu_id;type:int(11);" json:"menu_id"` // 菜单ID
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Uri      string `gorm:"colum:name;type:varchar(255);" json:"uri"`        // 请求地址
	Sort     int    `gorm:"column:sort;type:int(11);default:0;" json:"sort"` // 排序规则，默认升序，值越少越靠前
}

func (table *FunctionBasic) TableName() string {
	return "function_basic"
}

// GetFunctionList 获取功能列表
func GetFunctionList() *gorm.DB {
	return DB.Model(new(FunctionBasic)).Select("function_basic.identity, mb.identity menu_identity, " +
		"function_basic.name, function_basic.uri, function_basic.sort").
		Joins("LEFT JOIN menu_basic mb ON mb.id = function_basic.menu_id").
		Order("function_basic.sort ASC")
}
