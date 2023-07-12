package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	facade2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/lib/pq"
	"time"
)

var _ facade.IServerService = new(ServerService)

type ServerService struct {
	ServerRepo       facade2.IServerRepo       `inject:"ServerRepo"`
	ChannelRepo      facade2.IChannelRepo      `inject:"ChannelRepo"`
	ServerMemberRepo facade2.IServerMemberRepo `inject:"ServerMemberRepo"`
	IdentityRepo     facade2.IIdentityRepo     `inject:"IdentityRepo"`
	LabelRepo        facade2.ILabelRepo        `inject:"LabelRepo"`
	MemberRoleRepo   facade2.IMemberRoleRepo   `inject:"MemberRoleRepo"`
	MessageRepo      facade2.IMessageRepo      `inject:"MessageRepo"`
	UserRepo         user.IUserRepo            `inject:"UserRepo"`
}

func (s *ServerService) FindAllChannelByServerId(serverId int64) ([]*models.Channel, error) {
	return s.ChannelRepo.FindAllByServerId(serverId)
}

func (s *ServerService) FindAllServerByUserEmail(userEmail string) ([]*models.Server, error) {
	return s.ServerRepo.FindAllServerByUser(userEmail)
}

// SaveServer 保存服务器
func (s *ServerService) SaveServer(server *models.Server, email string) error {
	if flag, err := s.ServerRepo.ExistServerInNameAndOwner(server.Name, email); flag != false || err != nil {
		return err
	}

	owner, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	session := orm.NewXorm().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err := session.Rollback()
			if err != nil {
				return
			}
		}
	}()
	var labels = make([]string, len(server.Labels))
	for i, label := range server.Labels {
		labels[i] = label
	}
	serverId, err := s.ServerRepo.SaveServer(session, &models.Server{
		Name:       server.Name,
		OwnerId:    owner.ID,
		OwnerEmail: email,
		CreateAt:   time.Now(),
	})
	if err != nil {
		return err
	}
	var v = make([]models.Label, 0)
	for _, label := range labels {
		exist, err := session.Table(constant.Label).Where("name = ?", label).Exist(&models.Label{})
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		v = append(v, models.Label{
			ServerId: serverId,
			Name:     label,
		})
	}
	if len(v) > 0 {
		if _, err = session.Insert(&v); err != nil {
			return err
		}
	}

	sql := `UPDATE server SET labels = $1 and update_at = $2 WHERE id = $3`

	if _, err = session.Exec(sql, pq.Array(labels), time.Now(), serverId); err != nil {
		return err
	}

	return session.Commit()
}
