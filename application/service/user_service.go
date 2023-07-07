package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/common"
	constant2 "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	vo2 "Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
	errors "errors"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/openpgp/errors"
	"time"
)

var _ facade.IUserService = new(UserService)

type UserService struct {
	UserRepo user.IUserRepo `inject:"UserRepo"`
}

// Login return a loginResponse from given loginDto
func (u *UserService) Login(dto *vo2.LoginDto) (*vo2.LoginResponse, error) {
	// 查找用户
	user, err := u.UserRepo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	var roleMap = make(map[string]any)
	if err = json.Unmarshal(user.Role, &roleMap); err != nil {
		return nil, err
	}
	var roleName string
	for k := range roleMap {
		roleName = k
		if k != "SUPER" {
			// 检查密码
			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
				return nil, err
			}
		}
	}
	// 创建新token
	token := common.CreateNewToken(user.Email, roleName, false)
	return &vo2.LoginResponse{
		Authentication: constant2.Header + token,
	}, nil
}

// Register return a registerResponse from given registerRequest
func (u *UserService) Register(request *vo2.RegisterRequest) (*vo2.RegisterResponse, error) {
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
	role := constant2.User()
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
	token := common.CreateNewToken(request.Email, "USER", false)
	if err != nil {
		return nil, err
	}
	return &vo2.RegisterResponse{
		Message:        "注册成功",
		Authentication: constant2.Header + token,
	}, nil
}
