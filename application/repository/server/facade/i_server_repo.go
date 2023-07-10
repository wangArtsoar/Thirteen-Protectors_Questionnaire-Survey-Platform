package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"github.com/go-xorm/xorm"
)

type IServerRepo interface {
	SaveServer(session *xorm.Session, server *models.Server) (int64, error)
	FindAllServerByUser(userEmail string) ([]*models.Server, error)
	FindServerByName(serverName string) ([]*models.Server, error)
	EditServerById(Id int64, server *models.Server) (int64, error)
}
