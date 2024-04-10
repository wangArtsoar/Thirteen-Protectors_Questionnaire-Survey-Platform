package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	pl "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
	"github.com/go-xorm/xorm"
)

type IServerRepo interface {
	SaveServer(session *xorm.Session, server *models.Server) (int64, error)
	FindAllServerByUser(userEmail string) ([]*models.Server, error)
	FindServerByName(serverName string) ([]*models.Server, error)
	EditServerById(session *xorm.Session, Id int64, server *models.Server) error
	ExistServerInNameAndOwner(serverName string, ownerEmail string) (bool, error)
	FindServerInIds(ids []int64, request pl.PageRequest) (pl.PageList[models.Server], error)
}
