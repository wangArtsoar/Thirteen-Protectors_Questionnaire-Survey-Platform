package models

type Token struct {
	Id          uint   `json:"id" xorm:"pk autoincr"`
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	IsValid     int    `json:"is_valid"`
}
