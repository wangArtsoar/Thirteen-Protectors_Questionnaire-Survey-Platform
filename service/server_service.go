package service

import "Thirteen-Protectors_Questionnaire-Survey-Platform/service/facade"

var _ facade.IServerService = new(ServerService)

type ServerService struct {
}
