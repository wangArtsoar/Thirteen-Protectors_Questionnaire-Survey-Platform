package vo

// LoginResponse 登录响应体
type LoginResponse struct {
	Authentication string `json:"authentication"`
}

// RegisterResponse 注册响应体
type RegisterResponse struct {
	Message        string `json:"message"`
	Authentication string `json:"authentication"`
}
