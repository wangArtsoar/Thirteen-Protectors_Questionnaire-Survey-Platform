package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/common"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/const"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/vo"
	errors "errors"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/openpgp/errors"
	"time"
)

var _ facade.IUserService = new(UserService)

type UserService struct {
	UserRepo repository.IUserRepo `inject:"UserRepo"`
}

// Login return a loginResponse from given loginDto
func (u *UserService) Login(dto *vo.LoginDto) (*vo.LoginResponse, error) {
	// 查找用户
	user, err := u.UserRepo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	var roleMap = make(map[string]any)
	if err = json.Unmarshal(user.Role, &roleMap); err != nil {
		return nil, err
	}
	for k := range roleMap {
		if k != "SUPER" {
			// 检查密码
			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
				return nil, err
			}
		}
	}
	// 创建新token
	token := common.CreateNewToken(user.Email, roleMap, false)
	return &vo.LoginResponse{
		Authentication: _const.Header + token,
	}, nil
}

// Register return a registerResponse from given registerRequest
func (u *UserService) Register(request *vo.RegisterRequest) (*vo.RegisterResponse, error) {
	// check email database if exist
	flag, err := u.UserRepo.ExistByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if flag {
		return nil, errors.New("the user has existed")
	}
	id := uuid.New().String()
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}
	role := _const.User()
	roleJSON, err := json.Marshal(role)
	_, err = u.UserRepo.SaveUser(models.User{
		ID:       id,
		Name:     request.Name,
		Email:    request.Email,
		Password: string(password),
		CreateAt: time.Now(),
		Role:     roleJSON,
	})
	if err != nil {
		return nil, err
	}
	token := common.CreateNewToken(request.Email, map[string]any{}, false)
	if err != nil {
		return nil, err
	}
	return &vo.RegisterResponse{
		Message:        "注册成功",
		Authentication: _const.Header + token,
	}, nil
}
