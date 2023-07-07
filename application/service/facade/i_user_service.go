package facade

import (
	vo2 "Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
)

type IUserService interface {
	Login(dto *vo2.LoginDto) (*vo2.LoginResponse, error)
	Register(request *vo2.RegisterRequest) (*vo2.RegisterResponse, error)
}
