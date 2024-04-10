package service

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/facade"

var _ facade.IAppServerSrv = new(AppServerSrv)

type AppServerSrv struct {
}
