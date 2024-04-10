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

// ServerRequest 创建服务器请求体
type ServerRequest struct {
	Name   string   `json:"name"`
	Labels []string `json:"labels"`
}

// ChannelRequest 创建频道请求体
type ChannelRequest struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

// ServerMemberRequest 创建成员请求体
type ServerMemberRequest struct {
	ServerID   int64  `json:"server_id"`
	MemberName string `json:"member_name"`
	InviteId   int64  `json:"invite_id"`
}

// IdentityRequest 创建身份组请求体
type IdentityRequest struct {
	Name       string `json:"name"`
	ServerID   int64  `json:"server_id"`
	MemberRole string `json:"member_role"`
}

// MemberRoleRequest 创建成员角色请求体
type MemberRoleRequest struct {
	Name        string   `json:"name"`
	ServerID    int64    `json:"server_id"`
	Permissions []string `json:"permissions"`
}

// MessageRequest 创建信息请求体
type MessageRequest struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}
