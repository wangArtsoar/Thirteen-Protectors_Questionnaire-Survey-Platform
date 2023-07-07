package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.IServerMemberRepo = new(ServerMemberRepo)

type ServerMemberRepo struct {
}
