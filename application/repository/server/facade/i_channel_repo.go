package facade

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"

type IChannelRepo interface {
	CreateChannel(channel *models.Channel) (int64, error)
	FindAllByServerId(serverId int64) ([]*models.Channel, error)
	FindOneByChannelName(channelName string) (*models.Channel, error)
}
