package vo

import "time"

// LoginResponse 登录响应体
type LoginResponse struct {
	Authentication string `json:"authentication"`
}

// RegisterResponse 注册响应体
type RegisterResponse struct {
	Message        string `json:"message"`
	Authentication string `json:"authentication"`
}

// MessageResponse 信息响应体
type MessageResponse struct {
	Content  string    `json:"content"`
	SendDate time.Time `json:"send_date"`
	SendName string    `json:"send_name"`
	SendId   int64     `json:"send_id"`
}
