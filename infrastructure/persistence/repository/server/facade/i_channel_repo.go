package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"github.com/go-xorm/xorm"
)

type IChannelRepo interface {
	CreateChannel(*xorm.Session, *models.Channel) (int64, error)
	FindAllByServerId(int, int64) ([]*models.Channel, error)
	FindOneByChannelName(string) (*models.Channel, error)
}
