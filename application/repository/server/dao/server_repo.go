package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/go-xorm/xorm"
)

var _ facade.IServerRepo = new(ServerRepo)

type ServerRepo struct {
}

func (s *ServerRepo) SaveServer(session *xorm.Session, server *models.Server) (int64, error) {
	return session.InsertOne(server)
}

func (s *ServerRepo) FindAllServerByUser(userEmail string) ([]*models.Server, error) {
	var servers []*models.Server
	return servers, orm.NewXorm().
		Where("owner_email = ? and is_delete = ?", userEmail, constant.Default).
		Find(servers)
}

func (s *ServerRepo) FindServerByName(serverName string) ([]*models.Server, error) {
	var servers []*models.Server
	return servers, orm.NewXorm().Where("name = ? and is_delete = ?", serverName, constant.Default).Find(servers)
}

func (s *ServerRepo) EditServerById(Id int64, server *models.Server) (int64, error) {
	return orm.NewXorm().Id(Id).Where("is_delete = ?", constant.Default).Update(server)
}
