package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	facade2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/goccy/go-json"
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

func (s *ServerService) FindMessageByLimit(limit int) ([]*models.Message, error) {
	return s.MessageRepo.FindLastRecords(limit)
}

func (s *ServerService) FindMessageByKeyword(keyword string) ([]*models.Message, error) {
	return s.MessageRepo.FindByKeywords(keyword)
}

// SaveMessage 保存信息
func (s *ServerService) SaveMessage(message *models.Message, userEmail string) (*models.Message, error) {
	member, err := s.ServerMemberRepo.FindByUser(userEmail)
	if err != nil {
		return nil, err
	}
	message = settingMessage(message, member)
	return s.MessageRepo.SaveMessage(message)
}

func settingMessage(message *models.Message, member *models.ServerMember) *models.Message {
	message.SenderId = member.Id
	message.SendName = member.MemberName
	return message
}

// SaveMemberRole 保存成员角色
func (s *ServerService) SaveMemberRole(role *models.MemberRole) error {
	if _, err := s.MemberRoleRepo.NewAMemberRole(orm.NewXorm().NewSession(), role); err != nil {
		return err
	}
	return nil
}

// SaveIdentity 保存身份组
func (s *ServerService) SaveIdentity(identity *models.Identity) error {
	if _, err := s.IdentityRepo.SaveIdentity(orm.NewXorm().NewSession(), identity); err != nil {
		return err
	}
	return nil
}

// SaveServerMember 保存服务器成员
func (s *ServerService) SaveServerMember(serverMember *models.ServerMember) error {
	userInfo, err := s.UserRepo.FindByEmail(serverMember.UserEmail)
	if err != nil {
		return err
	}
	// init ServerMember
	newServerMember := createNewServerMember(userInfo, serverMember)
	// find channel limit 10 from serverId
	channels, err := s.ChannelRepo.FindAllByServerId(10, serverMember.ServerId)
	if err != nil {
		return err
	}
	for _, channel := range channels {
		newServerMember.Channels = append(serverMember.Channels, channel.Name)
	}
	if err = s.ServerMemberRepo.NewServerMember(orm.NewXorm().NewSession(), newServerMember); err != nil {
		return err
	}
	return nil
}

// createNewServerMember
func createNewServerMember(userInfo *models.User, member *models.ServerMember) *models.ServerMember {
	return &models.ServerMember{
		ServerId:   member.ServerId,
		MemberName: member.MemberName,
		UserName:   userInfo.Name,
		UserId:     userInfo.ID,
		UserEmail:  member.UserEmail,
		InviteId:   member.InviteId,
		CreateAt:   time.Now(),
		IdentityId: member.IdentityId,
	}
}

// SaveChannel 保存频道
func (s *ServerService) SaveChannel(channel *models.Channel) error {
	var (
		err  error
		flag bool
	)
	session := orm.NewXorm().NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	flag, err = s.LabelRepo.ExistByName(channel.Label)
	if err != nil {
		return err
	}
	if !flag {
		var labels = []*models.Label{ToLabelModel(channel.ServerId, channel.Label)}
		if _, err = s.LabelRepo.SaveLabel(session, labels); err != nil {
			return err
		}
	}
	if _, err = s.ChannelRepo.CreateChannel(session, channel); err != nil {
		return err
	}
	return session.Commit()
}

// FindAllChannelByServerId 根据服务器ID查询全部频道
func (s *ServerService) FindAllChannelByServerId(serverId int64) ([]*models.Channel, error) {
	return s.ChannelRepo.FindAllByServerId(constant.No, serverId)
}

// FindAllServerByUserEmail 根据创建者查询全部服务器
func (s *ServerService) FindAllServerByUserEmail(userEmail string) ([]*models.Server, error) {
	return s.ServerRepo.FindAllServerByUser(userEmail)
}

// SaveServer 保存服务器
func (s *ServerService) SaveServer(server *models.Server, email string) error {
	// check server
	if flag, err := s.ServerRepo.ExistServerInNameAndOwner(server.Name, email); flag != false || err != nil {
		return err
	}
	// get owner information
	owner, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	// begin a database session
	session := orm.NewXorm().NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		return err
	}
	// save server
	var serverId int64
	if serverId, err = s.ServerRepo.SaveServer(session, ToServerModel(owner.ID, owner.Email, server.Name)); err != nil {
		return err
	}
	var v = make([]*models.Label, 0)
	for _, label := range server.Labels {
		var exist bool
		if exist, err = s.LabelRepo.ExistByName(label); err != nil {
			session.Rollback()
			return err
		}
		if exist {
			continue
		}
		v = append(v, ToLabelModel(serverId, label))
	}
	// save label
	if len(v) > 0 {
		if _, err = session.Insert(v); err != nil {
			session.Rollback()
			return err
		}
	}
	// update server
	if err = s.ServerRepo.EditServerById(session, serverId, server); err != nil {
		session.Rollback()
		return err
	}
	// save server identity
	var identityID int64
	if identityID, err = s.IdentityRepo.SaveIdentity(session, ToIdentityModel(serverId)); err != nil {
		session.Rollback()
		return err
	}
	// join in server member
	if err = s.ServerMemberRepo.NewServerMember(session, ToServerMemberModel(serverId, identityID, owner)); err != nil {
		session.Rollback()
		return err
	}
	// save server member role
	if _, err = s.MemberRoleRepo.NewAMemberRole(session, ToMemberRoleModel(serverId)); err != nil {
		session.Rollback()
		return err
	}
	// update user serverIds
	owner.ServerIds, _ = json.Marshal(serverId)
	if _, err = session.Exec(`UPDATE "user" SET server_ids = ? WHERE email = ?`,
		owner.ServerIds, owner.Email); err != nil {
		return err
	}
	// commit
	return session.Commit()
}
