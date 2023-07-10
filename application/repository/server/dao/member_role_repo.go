package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
)

var _ facade.IMemberRoleRepo = new(MemberRoleRepo)

type MemberRoleRepo struct {
}

func (m *MemberRoleRepo) NewAMemberRole(role *models.MemberRole) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemberRoleRepo) UpdateMemberRole(memberRoleId int64, role *models.MemberRole) (int64, error) {
	//TODO implement me
	panic("implement me")
}
