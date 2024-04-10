package user

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"github.com/go-xorm/xorm"
)

type IUserRepo interface {
	ExistByEmail(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
	SaveUser(session *xorm.Session, user *models.User, serverId string) (int64, error)
}
