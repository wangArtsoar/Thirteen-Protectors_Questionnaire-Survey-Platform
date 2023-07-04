package repository

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/orm"
)

var _ IUserRepo = new(UserRepo)

type UserRepo struct {
}

func (u *UserRepo) SaveUser(user models.User) (int64, error) {
	return orm.NewXorm().InsertOne(user)
}

func (u *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := orm.NewXorm().Find(&user, "email = ?", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) ExistByEmail(email string) (bool, error) {
	return orm.NewXorm().Where("email = ?", email).Exist(&models.User{})
}
