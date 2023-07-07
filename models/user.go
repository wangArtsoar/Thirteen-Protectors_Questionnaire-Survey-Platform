package models

import (
	"time"
)

type User struct {
	ID        string    `json:"ID" xorm:"'id'"`
	Role      []byte    `json:"role"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ServerIds []byte    `json:"server_ids"` // 服务器列表
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"update_at"`
	IsDelete  int       `json:"is_delete"` // 是否删除(0-否,1-是)
	IsValid   int       `json:"is_valid"`  // 是否失效(0-否,1-是)
}
