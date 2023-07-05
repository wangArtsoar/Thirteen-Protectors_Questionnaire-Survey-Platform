package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/bean"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/common"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/vo"
	errors "errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/openpgp/errors"
	"time"
)

var _ facade.IUserService = new(UserService)

type UserService struct {
	UserRepo  repository.IUserRepo  `inject:"UserRepo"`
	RoleRepo  repository.IRoleRepo  `inject:"RoleRepo"`
	TokenRepo repository.ITokenRepo `inject:"TokenRepo"`
}

// Login return a loginResponse from given loginDto
func (u UserService) Login(dto *vo.LoginDto) (*vo.LoginResponse, error) {
	// 查找用户
	user, err := u.UserRepo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	// 检查密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, err
	}
	// 删除token
	if _, err = u.TokenRepo.DeleteTokenByUserId(user.ID); err != nil {
		return nil, err
	}
	// 创建新token
	token := common.CreateNewToken(user.Email, false)
	// 保存token
	if _, err = u.TokenRepo.SaveToken(models.Token{UserId: user.ID, AccessToken: token}); err != nil {
		return nil, err
	}
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
	roleId, err := u.RoleRepo.GetIdByRoleName(request.RoleName)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}
	_, err = u.UserRepo.SaveUser(models.User{
		ID:       id,
		Name:     request.Name,
		Email:    request.Email,
		Password: string(password),
		CreateAt: time.Now(),
		RoleId:   roleId,
	})
	if err != nil {
		return nil, err
	}
	token := common.CreateNewToken(request.Email, false)
	_, err = u.TokenRepo.SaveToken(models.Token{
		AccessToken: token,
		UserId:      id,
	})
	if err != nil {
		return nil, err
	}
	return &vo.RegisterResponse{
		Message:        "注册成功",
		Authentication: bean.Header + token,
	}, nil
}
