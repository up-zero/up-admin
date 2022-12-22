package login

type PasswordRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type PasswordReply struct {
	Token        string `json:"token"`         // token
	RefreshToken string `json:"refresh_token"` // 用于刷新 token 的 token
}
