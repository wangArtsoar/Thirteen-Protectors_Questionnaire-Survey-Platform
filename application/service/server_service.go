package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	facade2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
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

func (s *ServerService) SaveServer(server *models.Server, email string) error {
	owner, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	server.OwnerId = owner.ID
	server.OwnerEmail = email

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

	labels := server.Labels
	_, err = session.Insert(&labels)
	if err != nil {
		return err
	}

	_, err = s.ServerRepo.SaveServer(session, server)
	if err != nil {
		return err
	}
	return session.Commit()
}
