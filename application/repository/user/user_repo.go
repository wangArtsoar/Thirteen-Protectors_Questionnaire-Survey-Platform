package user

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/go-xorm/xorm"
)

var _ IUserRepo = new(UserRepo)

type UserRepo struct {
}

func (u *UserRepo) SaveUser(session *xorm.Session, user models.User) (int64, error) {
	return session.InsertOne(user)
}

func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if _, err := orm.NewXorm().Where(
		"email = ? and is_delete = ? and is_valid = ?", email, constant.No, constant.No).
		Get(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) ExistByEmail(email string) (bool, error) {
	return orm.NewXorm().Where(
		"email = ? and is_delete = ? and is_valid = ?", email, constant.No, constant.No).
		Exist(&models.User{})
}
