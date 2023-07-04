package repository

import "Thirteen-Protectors_Questionnaire-Survey-Platform/models"

type IRoleRepo interface {
	SaveRole(role models.Role) (int64, error)
	ExistRole(roleName string) (bool, error)
	GetIdByRoleName(roleName string) (uint, error)
}
