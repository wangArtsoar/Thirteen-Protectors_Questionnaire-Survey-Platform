package user

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"github.com/go-xorm/xorm"
)

var _ IUserRepo = new(UserRepo)

type UserRepo struct {
}

// NewUserRepo 构造器
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) SaveUser(session *xorm.Session, user *models.User, serverId string) (int64, error) {
	if len(user.ID) != 0 {
		return updateUser(session, user, serverId)
	}
	return session.InsertOne(user)
}

func updateUser(session *xorm.Session, user *models.User, serverId string) (int64, error) {
	sql := `UPDATE "user" SET server_ids = jsonb_insert(server_ids::jsonb,'{-1}',?::jsonb,true) WHERE email = ?`
	if _, err := session.Exec(sql, serverId, user.Email); err != nil {
		return 0, err
	}
	return 1, nil
}

func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if _, err := persistence.NewXorm().Where(
		"email = ? and is_delete = ? and is_valid = ?", email, constant.No, constant.No).
		Get(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) ExistByEmail(email string) (bool, error) {
	return persistence.NewXorm().Where(
		"email = ? and is_delete = ? and is_valid = ?", email, constant.No, constant.No).
		Exist(&models.User{})
}
