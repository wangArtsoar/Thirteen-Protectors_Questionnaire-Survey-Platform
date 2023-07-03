package service

import "Thirteen-Protectors_Questionnaire-Survey-Platform/vo"

type IUserService interface {
	Login(dto *vo.LoginDto) (*vo.LoginResponse, error)
	Register(request *vo.RegisterRequest) (*vo.RegisterResponse, error)
}
