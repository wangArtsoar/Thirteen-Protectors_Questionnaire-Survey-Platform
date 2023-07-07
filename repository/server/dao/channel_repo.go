package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.IChannelRepo = new(ChannelRepo)

type ChannelRepo struct {
}
