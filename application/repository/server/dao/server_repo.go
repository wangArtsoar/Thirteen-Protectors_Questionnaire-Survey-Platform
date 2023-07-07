package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
)

var _ facade.IServerRepo = new(ServerRepo)

type ServerRepo struct {
}
