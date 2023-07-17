package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"github.com/go-xorm/xorm"
)

type IMemberRoleRepo interface {
	NewAMemberRole(session *xorm.Session, role *models.MemberRole) (int64, error)
	UpdateMemberRole(memberRoleId int64, role *models.MemberRole) (int64, error)
}
