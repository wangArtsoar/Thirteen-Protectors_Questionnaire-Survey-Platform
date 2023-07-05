package models

import (
	"time"
)

type User struct {
	ID       string    `json:"ID" xorm:"'id'"`
	RoleId   uint      `json:"role_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"update_at"`
	IsDelete int       `json:"is_delete"`
	IsValid  int       `json:"is_valid"`
}

type Role struct {
	Id   uint   `json:"id" xorm:"pk autoincr"`
	Name string `json:"name"`
}
