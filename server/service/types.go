package service

// ------------- COMMON -------------

type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	KeyWord string `json:"key_word" form:"key_word"`
}

// ------------- COMMON -------------

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

type UserInfoReply struct {
	Username string `json:"username"`  // 用户名
	Phone    string `json:"phone"`     // 手机号
	Avatar   string `json:"avatar"`    // 头像
	RoleName string `json:"role_name"` // 角色名称
}

type UserPasswordChangeRequest struct {
	OldPassword string `json:"old_password"` // 旧密码
	NewPassword string `json:"new_password"` // 新密码
}

type SetRoleListRequest struct {
	*QueryRequest
}

type SetRoleListReply struct {
	Identity  string `json:"identity"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	IsAdmin   int    `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
