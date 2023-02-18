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
	WebIcon  string `json:"web_icon"`
	Sort     int    `json:"sort"`
	Path     string `json:"path"`
}

type MenuReply struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	WebIcon  string `json:"web_icon"`
	Sort     int    `json:"sort"`
	Path     string `json:"path"`
	SubMenus []struct {
		Identity string `json:"identity"`
		Name     string `json:"name"`
		WebIcon  string `json:"web_icon"`
		Sort     int    `json:"sort"`
		Path     string `json:"path"`
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

type SetRoleUpdateAdminRequest struct {
	Identity string `json:"identity"` // 角色唯一标识
	IsAdmin  int    `json:"is_admin"` // 是否是超管【0-否 1-是】
}

type SetRoleCreateRequest struct {
	Name           string   `json:"name"`
	Sort           int64    `json:"sort"`
	IsAdmin        int8     `json:"is_admin"`
	MenuIdentities []string `json:"menu_identities"` // 被授权的菜单列表
	FuncIdentities []string `json:"func_identities"` // 被授权的功能列表
}

type SetRoleUpdateRequest struct {
	Identity string `json:"identity"` // 角色唯一标识
	SetRoleCreateRequest
}

type SetFuncListReply struct {
	Identity     string `json:"identity"`
	MenuIdentity string `json:"menu_identity"`
	Name         string `json:"name"`
	Uri          string `json:"uri"`
	Sort         int    `json:"sort"`
}

type SetRoleDetailReply struct {
	SetRoleCreateRequest
}

type DevMenuAddRequest struct {
	ParentIdentity string `json:"parent_identity"` // 父级唯一标识，不填默认为顶级菜单
	Name           string `json:"name"`            // 菜单名称
	WebIcon        string `json:"web_icon"`        // 网页图标
	Sort           int    `json:"sort"`            // 排序
}

type DevMenuUpdateRequest struct {
	Identity string `json:"identity"` // 菜单唯一标识，必填
	DevMenuAddRequest
}

type DevFuncAddRequest struct {
	MenuIdentity string `json:"menu_identity"` // 菜单唯一标识
	Name         string `json:"name"`          // 功能名称
	Uri          string `json:"uri"`           // 请求地址
	Sort         int    `json:"sort"`          // 排序
}

type DevFuncUpdateRequest struct {
	Identity string `json:"identity"` // 功能唯一标识，必填
	DevFuncAddRequest
}
