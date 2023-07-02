package models

import "time"

type User struct {
	UUID     string    `json:"UUID"`
	RoleId   uint      `json:"role_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"update_at"`
	IsDelete bool      `json:"is_delete"`
	IsValid  bool      `json:"is_valid"`
}

type Role struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
