package models

import "gorm.io/gorm"

type MenuBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ParentId uint   `gorm:"column:parent_id;type:int(11);" json:"parent_id"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Sort     int    `gorm:"column:sort;type:int(11);default:0;" json:"sort"` // 排序规则，默认升序，值越少越靠前
}

func (table *MenuBasic) TableName() string {
	return "menu_basic"
}
