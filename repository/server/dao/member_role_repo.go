package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.IMemberRoleRepo = new(MemberRoleRepo)

type MemberRoleRepo struct {
}