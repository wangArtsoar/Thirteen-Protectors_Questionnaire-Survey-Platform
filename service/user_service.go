package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/bean"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/common"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/vo"
	errors "errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/openpgp/errors"
	"time"
)

var _ IUserService = new(UserService)

type UserService struct {
	UserRepo repository.IUserRepo `inject:"userRepo"`
}

// Login return a loginResponse from given loginDto
func (u UserService) Login(dto *vo.LoginDto) (*vo.LoginResponse, error) {
	user, err := u.UserRepo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, err
	}
	token := common.CreateNewToken(user.Email)
	return &vo.LoginResponse{
		Authentication: bean.Header + token,
	}, nil
}

// Register return a registerResponse from given registerRequest
func (u UserService) Register(request *vo.RegisterRequest) (*vo.RegisterResponse, error) {
	// check email database if exist
	flag, err := u.UserRepo.ExistByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if flag {
		return nil, errors.New("the user has existed")
	}
	roleId, err := u.UserRepo.GetIdByRoleName(request.RoleName)
	if err != nil {
		return nil, err
	}
	_, err = u.UserRepo.SaveUser(models.User{
		UUID:     uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		CreateAt: time.Now(),
		RoleId:   roleId,
	})
	if err != nil {
		return nil, err
	}
	token := common.CreateNewToken(request.Email)
	return &vo.RegisterResponse{
		Message:        "注册成功",
		Authentication: bean.Header + token,
	}, nil
}
