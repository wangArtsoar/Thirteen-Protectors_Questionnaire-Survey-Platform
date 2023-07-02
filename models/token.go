package models

type Token struct {
	Id          uint   `json:"id"`
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	IsValid     bool   `json:"is_valid"`
}
