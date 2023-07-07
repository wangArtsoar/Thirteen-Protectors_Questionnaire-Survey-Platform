package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.IServerRepo = new(ServerRepo)

type ServerRepo struct {
}
