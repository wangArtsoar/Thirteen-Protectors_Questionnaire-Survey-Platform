package vo

// LoginDto 登录请求体
type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest 注册请求体
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleName string `json:"roleName"`
}
