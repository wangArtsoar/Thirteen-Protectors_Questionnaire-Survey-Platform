package repository

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/orm"
)

var _ IRoleRepo = new(RoleRepo)

type RoleRepo struct {
}

func (r *RoleRepo) SaveRole(role models.Role) (int64, error) {
	return orm.NewXorm().InsertOne(role)
}

func (r *RoleRepo) ExistRole(roleName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RoleRepo) GetIdByRoleName(roleName string) (uint, error) {
	var role models.Role
	if err := orm.NewXorm().Cols("id").Where("name = ?", roleName).Find(&role); err != nil {
		return -1, err
	}
	return role.Id, nil
}
