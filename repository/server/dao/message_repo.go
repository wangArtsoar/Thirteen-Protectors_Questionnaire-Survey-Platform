package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.IMessageRepo = new(MessageRepo)

type MessageRepo struct {
}
