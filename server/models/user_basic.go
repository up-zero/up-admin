package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity  string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Username  string `gorm:"column:username;type:varchar(50);" json:"username"`
	Password  string `gorm:"column:password;type:varchar(36);" json:"password"`
	Phone     string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(255);" json:"wx_union_id"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(255);" json:"wx_open_id"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// GetUserBasicByUsernamePassword 根据用户名、密码查询 user_basic
func GetUserBasicByUsernamePassword(username, password string) (*UserBasic, error) {
	data := new(UserBasic)
	err := DB.Where("username = ? AND password = ?", username, password).First(data).Error
	return data, err
}
