package service

import (
	vo "Thirteen-Protectors_Questionnaire-Survey-Platform/application/vo"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/dao"
	repo "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/facade"
	user "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
	"errors"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/openpgp/errors"
	"time"
)

var _ IUserService = new(UserService)

type UserService struct {
	UserRepo   user.IUserRepo   `dig:"UserRepo"`
	ServerRepo repo.IServerRepo `dig:"ServerRepo"`
}

func NewUserService() IUserService {
	return &UserService{user.NewUserRepo(), dao.NewServerRepo()}
}

// Login return a loginResponse from given loginDto
func (srv *UserService) Login(dto *vo.LoginDto) (*vo.LoginResponse, error) {
	// 查找用户
	userInfo, err := srv.UserRepo.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}
	var roleMap = make(map[string]any)
	if err = json.Unmarshal(userInfo.Role, &roleMap); err != nil {
		return nil, err
	}
	var roleName string
	for k := range roleMap {
		roleName = k
		if k != "SUPER" {
			// 检查密码
			if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(dto.Password)); err != nil {
				return nil, err
			}
		}
	}
	// 创建新token
	token := util.CreateNewToken(userInfo.Email, roleName, false)
	return &vo.LoginResponse{
		Authentication: constant.Header + token,
	}, nil
}

// Register return a registerResponse from given registerRequest
func (srv *UserService) Register(request *vo.RegisterRequest) (*vo.RegisterResponse, error) {
	if err := checkEmail(srv, request.Email); err != nil {
		return nil, err
	}
	// create a new User struct
	_, newUser, err := createNewUser(request)
	// begin a database session
	session := persistence.NewXorm().NewSession()
	if err = session.Begin(); err != nil {
		return nil, err
	}
	defer session.Close()
	// save user
	if _, err = srv.UserRepo.SaveUser(session, &newUser, ""); err != nil {
		return nil, err
	}
	//// save server
	//var serverId int64
	//if serverId, err = srv.ServerRepo.SaveServer(session, createNewServer(id, request.Email, "")); err != nil {
	//	return nil, err
	//}
	//// update user
	//newUser.ServerIds, _ = json.Marshal(serverId)
	//if _, err = session.Table(constant.UserTable).Where("email = ?", request.Email).Update(&newUser); err != nil {
	//	return nil, err
	//}
	// commit
	if err = session.Commit(); err != nil {
		return nil, err
	}
	token := util.CreateNewToken(request.Email, "USER", false)
	if err != nil {
		return nil, err
	}
	return &vo.RegisterResponse{
		Message:        "注册成功",
		Authentication: constant.Header + token,
	}, nil
}

// createNewUser
func createNewUser(request *vo.RegisterRequest) (string, models.User, error) {
	id := uuid.New().String()
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return "", models.User{}, err
	}
	role := constant.User()
	roleJSON, err := json.Marshal(role)
	serverIds, err := json.Marshal([]string{})
	return id, models.User{
		ID:        id,
		Name:      request.Name,
		Email:     request.Email,
		Password:  string(password),
		CreateAt:  time.Now(),
		Role:      roleJSON,
		ServerIds: serverIds,
	}, nil
}

// checkEmail
func checkEmail(srv *UserService, email string) error {
	// check email database if exist
	flag, err := srv.UserRepo.ExistByEmail(email)
	if err != nil {
		return err
	}
	if flag {
		return errors.New("the user has existed")
	}
	return nil
}
