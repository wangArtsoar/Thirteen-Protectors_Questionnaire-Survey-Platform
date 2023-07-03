package repository

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/orm"
)

var _ IUserRepo = new(UserRepo)

type UserRepo struct {
}

func (u UserRepo) SaveRole(role models.Role) (int64, error) {
	one, err := orm.NewXorm().InsertOne(role)
	if err != nil {
		return 0, err
	}
	return one, nil
}

func (u UserRepo) GetIdByRoleName(roleName string) (uint, error) {
	var role models.Role
	err := orm.NewXorm().Cols("").Find(&role)
	if err != nil {
		return 0, err
	}
	return role.Id, nil
}

func (u UserRepo) SaveUser(user models.User) (int64, error) {
	id, err := orm.NewXorm().InsertOne(user)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (u UserRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := orm.NewXorm().Find(&user, "email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u UserRepo) ExistByEmail(email string) (bool, error) {
	exist, err := orm.NewXorm().Exist(email)
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}
	return false, nil
}
