package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	pl "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/page_list"
	"github.com/go-xorm/xorm"
	"github.com/lib/pq"
	"time"
)

var _ facade.IServerRepo = new(ServerRepo)

type ServerRepo struct {
}

func (s *ServerRepo) FindServerInIds(ids []int64, page pl.PageRequest) (pl.PageList[models.Server], error) {
	var servers []models.Server
	err := orm.NewXorm().In("id", ids).Limit(page.PageSize, page.PageNum-1).Find(&servers)
	if err != nil {
		return pl.PageList[models.Server]{}, err
	}
	return pl.Pageable(servers, page, len(ids)), nil
}

// NewServerRepo 构造器
func NewServerRepo() *ServerRepo {
	return &ServerRepo{}
}

func (s *ServerRepo) ExistServerInNameAndOwner(serverName string, ownerEmail string) (bool, error) {
	return orm.NewXorm().Where("Name = ? and owner_email = ?", serverName, ownerEmail).Exist(&models.Server{})
}

func (s *ServerRepo) SaveServer(session *xorm.Session, server *models.Server) (int64, error) {
	var lastInsertId int64
	sql := `INSERT INTO server(name, create_at, owner_id, owner_email) VALUES (?, ?, ?, ?) RETURNING id`
	_, err := session.SQL(sql, server.Name, server.CreateAt, server.OwnerId, server.OwnerEmail).Get(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (s *ServerRepo) FindAllServerByUser(userEmail string) ([]*models.Server, error) {
	var servers []*models.Server
	return servers, orm.NewXorm().
		Where("owner_email = ? and is_delete = ?", userEmail, constant.No).
		Find(servers)
}

func (s *ServerRepo) FindServerByName(serverName string) ([]*models.Server, error) {
	var servers []*models.Server
	return servers, orm.NewXorm().Where("name = ? and is_delete = ?", serverName, constant.No).Find(servers)
}

func (s *ServerRepo) EditServerById(session *xorm.Session, Id int64, server *models.Server) error {
	sql := `UPDATE server SET labels = ?,update_at = ? WHERE id = ?`
	labels := server.Labels
	if _, err := session.Exec(sql, pq.Array(labels), time.Now(), Id); err != nil {
		return err
	}
	return nil
}
