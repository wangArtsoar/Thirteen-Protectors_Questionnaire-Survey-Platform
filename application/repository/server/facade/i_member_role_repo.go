package facade

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"

type IMemberRoleRepo interface {
	NewAMemberRole(role *models.MemberRole) (int64, error)
	UpdateMemberRole(memberRoleId int64, role *models.MemberRole) (int64, error)
}
