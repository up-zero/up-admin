package service

type LoginPasswordRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type LoginPasswordReply struct {
	Token        string `json:"token"`         // token
	RefreshToken string `json:"refresh_token"` // 用于刷新 token 的 token
}

type RoleMenu struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parent_id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
}

type MenuReply struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	SubMenus []struct {
		Identity string `json:"identity"`
		Name     string `json:"name"`
		Sort     int    `json:"sort"`
	} `json:"sub_menus"`
}
