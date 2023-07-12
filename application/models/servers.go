package models

import (
	"time"
)

// Server 服务器
type Server struct {
	Id         int64     `json:"id" xorm:"'id' pk autoincr"`
	Name       string    `json:"name"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	OwnerId    string    `json:"owner_id" xorm:"owner_id index"`
	OwnerEmail string    `json:"owner_email" xorm:"owner_email index"`
	Labels     []string  `json:"labels"`
	IsDelete   int       `json:"is_delete"`
}

// Channel 频道
type Channel struct {
	Id       int64     `json:"id" xorm:"'id' pk autoincr"`
	Label    string    `json:"label"`
	ServerId int64     `json:"server_id" xorm:"server_id index"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	IsMute   int       `json:"is_mute"`
	IsDelete int       `json:"is_delete"`
}

// ServerMember 服务器成员
type ServerMember struct {
	Id         int64     `json:"id" xorm:"'id' pk autoincr"`
	UserId     string    `json:"user_id" xorm:"user_id unique"`
	UserEmail  string    `json:"user_email"`
	UserName   string    `json:"user_name"`
	ServerId   int64     `json:"server_id"`
	MemberName string    `json:"member_name"`
	IdentityId int64     `json:"identity_id" xorm:"identity_id unique"`
	InviteId   int64     `json:"invite_id"` // 邀请人
	Channels   []string  `json:"channels"`  // 频道列表
	CreateAt   time.Time `json:"create_at"`
	IsDelete   int       `json:"is_delete"`
	IsRobot    int       `json:"is_robot"`   // 是否为机器人
	IsMute     int       `json:"is_mute"`    // 是否被禁言
	IsWaiting  int       `json:"is_waiting"` // 是否在等待通过
}

// Identity 身份
type Identity struct {
	Id         int64  `json:"id" xorm:"'id' pk autoincr"`
	ChannelId  int64  `json:"channel_id"`
	Name       string `json:"name" xorm:"name unique"`
	MemberRole string `json:"member_role"`
}

// MemberRole 成员角色
type MemberRole struct {
	Id          int64    `json:"id" xorm:"'id' pk autoincr"`
	ChannelId   int64    `json:"channel_id"`
	Name        string   `json:"name" xorm:"name unique"`
	Permissions []string `json:"permissions"`
}

// Label 标签
type Label struct {
	Id       int64  `json:"id" xorm:"'id' pk autoincr"`
	ServerId int64  `json:"server_id"`
	Name     string `json:"name" xorm:"name unique"`
}

// Message 消息
type Message struct {
	Id          int64     `json:"id" xorm:"'id' pk autoincr"`
	Message     string    `json:"message"`
	Image       string    `json:"image"`
	Video       string    `json:"video"`
	SenderId    int64     `json:"sender_id"` // ServerMember Id
	Date        time.Time `json:"date"`
	IsWithdraw  int       `json:"is_withdraw"`
	LimitedTime time.Time `json:"limited_time"` // 撤回限时
}
