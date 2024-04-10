package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"github.com/go-xorm/xorm"
)

type IIdentityRepo interface {
	SaveIdentity(session *xorm.Session, identity *models.Identity) (int64, error)
	FindAllIdentityByChannelId(channelId int64, err error) ([]*models.Identity, error)
}
