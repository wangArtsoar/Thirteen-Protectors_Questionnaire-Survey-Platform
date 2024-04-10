package service

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/facade"

var _ facade.IAppUserSrv = new(AppUserSrv)

type AppUserSrv struct {
}
