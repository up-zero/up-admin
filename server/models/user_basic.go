package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity     string    `gorm:"column:identity;type:varchar(36);" json:"identity"`
	RoleIdentity string    `gorm:"column:role_identity;type:varchar(36);" json:"role_identity"`     // 角色唯一标识
	RoleBasic    RoleBasic `gorm:"foreignKey:identity;references:role_identity;" json:"role_basic"` // 管理角色基础表
	Username     string    `gorm:"column:username;type:varchar(50);" json:"username"`
	Password     string    `gorm:"column:password;type:varchar(36);" json:"password"`
	Phone        string    `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId    string    `gorm:"column:wx_union_id;type:varchar(255);" json:"wx_union_id"`
	WxOpenId     string    `gorm:"column:wx_open_id;type:varchar(255);" json:"wx_open_id"`
	Avatar       string    `gorm:"column:avatar;type:varchar(255);" json:"avatar"` // 头像
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

// GetUserInfo 获取用户详情
func GetUserInfo(identity string) *gorm.DB {
	tx := DB.Debug().Model(new(UserBasic)).Select("user_basic.username, user_basic.phone, user_basic.avatar, rb.name role_name").
		Joins("LEFT JOIN role_basic rb ON rb.identity = user_basic.role_identity").
		Where("user_basic.identity = ?", identity)
	return tx
}

// GetUserList 获取管理员列表
func GetUserList(keyword string) *gorm.DB {
	tx := DB.Model(new(UserBasic)).Select("user_basic.identity, user_basic.role_identity, rb.name role_name, " +
		"user_basic.username, user_basic.phone, user_basic.created_at, user_basic.updated_at").
		Joins("LEFT JOIN role_basic rb ON rb.identity = user_basic.role_identity")
	if keyword != "" {
		tx.Where("user_basic.username LIKE ?", "%"+keyword+"%")
	}
	return tx
}
