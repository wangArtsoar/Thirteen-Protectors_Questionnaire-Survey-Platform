package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.IIdentityRepo = new(IdentityRepo)

type IdentityRepo struct {
}