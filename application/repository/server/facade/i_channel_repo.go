package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"github.com/go-xorm/xorm"
)

type IChannelRepo interface {
	CreateChannel(session *xorm.Session, channel *models.Channel) (int64, error)
	FindAllByServerId(limit int, serverId int64) ([]*models.Channel, error)
	FindOneByChannelName(channelName string) (*models.Channel, error)
}
