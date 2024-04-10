package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	model "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/dao"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/user"
	pl "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
	"github.com/goccy/go-json"
	"github.com/samber/lo"
	"strconv"
	"time"
)

var _ IServerService = new(ServerService)

type ServerService struct {
	ServerRepo       facade.IServerRepo       `dig:"method=dao.NewServerRepo"`
	ChannelRepo      facade.IChannelRepo      `dig:"method=dao.NewChannelRepo"`
	ServerMemberRepo facade.IServerMemberRepo `dig:"method=dao.NewServerMemberRepo"`
	IdentityRepo     facade.IIdentityRepo     `dig:"method=dao.NewIdentityRepo"`
	LabelRepo        facade.ILabelRepo        `dig:"method=dao.NewLabelRepo"`
	MemberRoleRepo   facade.IMemberRoleRepo   `dig:"method=dao.NewMemberRoleRepo"`
	MessageRepo      facade.IMessageRepo      `dig:"method=dao.NewMessageRepo"`
	UserRepo         user.IUserRepo           `dig:"method=dao.NewUserRepo"`
}

func (s *ServerService) FindJoinServerListByUser(userEmail string, page pl.PageRequest) (*pl.PageList[model.Server], error) {
	serverMembers, err := s.ServerMemberRepo.FindByUser(userEmail)
	if err != nil {
		return nil, err
	}
	serverIds := lo.Map(serverMembers, func(item model.ServerMember, index int) int64 {
		return item.ServerId
	})
	servers, err := s.ServerRepo.FindServerInIds(serverIds, page)
	if err != nil {
		return nil, err
	}
	return &servers, nil
}

func (s *ServerService) FindMessageByLimit(limit int) ([]*model.Message, error) {
	return s.MessageRepo.FindLastRecords(limit)
}

func (s *ServerService) FindMessageByKeyword(keyword string) ([]*model.Message, error) {
	return s.MessageRepo.FindByKeywords(keyword)
}

// SaveMessage 保存信息
func (s *ServerService) SaveMessage(message *model.Message, userEmail string) (*model.Message, error) {
	member, err := s.ServerMemberRepo.FindByUser(userEmail)
	if err != nil {
		return nil, err
	}
	message = settingMessage(message, &member[0])
	return s.MessageRepo.SaveMessage(message)
}

func settingMessage(message *model.Message, member *model.ServerMember) *model.Message {
	message.SenderId = member.Id
	message.SendName = member.MemberName
	return message
}

// SaveMemberRole 保存成员角色
func (s *ServerService) SaveMemberRole(role *model.MemberRole) error {
	if _, err := s.MemberRoleRepo.NewAMemberRole(persistence.NewXorm().NewSession(), role); err != nil {
		return err
	}
	return nil
}

// SaveIdentity 保存身份组
func (s *ServerService) SaveIdentity(identity *model.Identity) error {
	if _, err := s.IdentityRepo.SaveIdentity(persistence.NewXorm().NewSession(), identity); err != nil {
		return err
	}
	return nil
}

// SaveServerMember 保存服务器成员
func (s *ServerService) SaveServerMember(serverMember *model.ServerMember) error {
	userInfo, err := s.UserRepo.FindByEmail(serverMember.UserEmail)
	if err != nil {
		return err
	}
	// init ServerMember
	newServerMember := createNewServerMember(userInfo, serverMember)
	// find channel limit 10 from serverId
	var channels []*model.Channel
	if channels, err = s.ChannelRepo.FindAllByServerId(10, serverMember.ServerId); err != nil {
		return err
	}
	for _, channel := range channels {
		newServerMember.Channels = append(serverMember.Channels, channel.Name)
	}
	session := persistence.NewXorm().NewSession()
	if err = session.Begin(); err != nil {
		return err
	}
	defer session.Close()
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	// save serverMember
	if _, err = s.ServerMemberRepo.NewServerMember(session, newServerMember); err != nil {
		return err
	}
	// update user serversID
	if _, err = s.UserRepo.SaveUser(session, userInfo, strconv.FormatInt(serverMember.ServerId, 10)); err != nil {
		return err
	}
	return session.Commit()
}

// createNewServerMember
func createNewServerMember(userInfo *model.User, member *model.ServerMember) *model.ServerMember {
	return &model.ServerMember{
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
func (s *ServerService) SaveChannel(channel *model.Channel) error {
	var (
		err  error
		flag bool
	)
	session := persistence.NewXorm().NewSession()
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
		var labels = []*model.Label{ToLabelModel(channel.ServerId, channel.Label)}
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
func (s *ServerService) FindAllChannelByServerId(serverId int64) ([]*model.Channel, error) {
	return s.ChannelRepo.FindAllByServerId(constant.No, serverId)
}

// FindAllServerByUserEmail 根据创建者查询全部服务器
func (s *ServerService) FindAllServerByUserEmail(userEmail string) ([]*model.Server, error) {
	return s.ServerRepo.FindAllServerByUser(userEmail)
}

// SaveServer 保存服务器
func (s *ServerService) SaveServer(server *model.Server, email string) error {
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
	session := persistence.NewXorm().NewSession()
	if err = session.Begin(); err != nil {
		return err
	}
	defer session.Close()
	defer func() {
		if err != nil {
			session.Rollback()
		}
	}()
	// save server
	var serverId int64
	if serverId, err = s.ServerRepo.SaveServer(session, ToServerModel(owner.ID, owner.Email, server.Name)); err != nil {
		return err
	}
	var v = make([]*model.Label, 0)
	var labels []string
	json.Unmarshal(server.Labels, labels)
	for _, label := range labels {
		var exist bool
		if exist, err = s.LabelRepo.ExistByName(label); err != nil {
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
			return err
		}
	}
	// update server
	if err = s.ServerRepo.EditServerById(session, serverId, server); err != nil {
		return err
	}
	// save server identity
	var identityID int64
	if identityID, err = s.IdentityRepo.SaveIdentity(session, ToIdentityModel(serverId)); err != nil {
		return err
	}
	// join in server member
	if _, err = s.ServerMemberRepo.NewServerMember(session, ToServerMemberModel(serverId, identityID, owner)); err != nil {
		return err
	}
	// save server member role
	if _, err = s.MemberRoleRepo.NewAMemberRole(session, ToMemberRoleModel(serverId)); err != nil {
		return err
	}
	// update user serverIds
	owner.ServerIds, _ = json.Marshal(serverId)
	if _, err = s.UserRepo.SaveUser(session, owner, ""); err != nil {
		return err
	}
	// commit
	return session.Commit()
}

func NewServerService() IServerService {
	return &ServerService{
		ServerRepo:       dao.NewServerRepo(),
		ChannelRepo:      dao.NewChannelRepo(),
		ServerMemberRepo: dao.NewServerMemberRepo(),
		IdentityRepo:     dao.NewIdentityRepo(),
		LabelRepo:        dao.NewLabelRepo(),
		MemberRoleRepo:   dao.NewMemberRoleRepo(),
		MessageRepo:      dao.NewMessgeRepo(),
		UserRepo:         user.NewUserRepo()}
}
