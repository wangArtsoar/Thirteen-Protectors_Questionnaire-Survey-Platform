package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
)

var _ facade.IServerMemberRepo = new(ServerMemberRepo)

type ServerMemberRepo struct {
}

func (s *ServerMemberRepo) NewAServerMember(member *models.ServerMember) (int64, error) {
	return orm.NewXorm().InsertOne(member)
}

func (s *ServerMemberRepo) FindAllServerMember() ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, orm.NewXorm().Find(serverMembers)
}

func (s *ServerMemberRepo) FindAllServerMemberByServerId(serverId int64) ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, orm.NewXorm().Where("server_id = ?", serverId).Find(serverMembers)
}

func (s *ServerMemberRepo) FindAllServerMemberByChannel(channel *models.Channel) ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, orm.NewXorm().Where("channel = ?", channel.Name).Find(serverMembers)
}

func (s *ServerMemberRepo) FindServerMemberByIdentityId(identityId int64) ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, orm.NewXorm().Where("identity_id = ?", identityId).Find(serverMembers)
}

func (s *ServerMemberRepo) UpdateServerMemberByUserEmail(userEmail string, member *models.ServerMember) (int64, error) {
	return orm.NewXorm().Where("user_email = ?", userEmail).Update(member)
}

func (s *ServerMemberRepo) DeleteServerMemberById(id int64) (int64, error) {
	return orm.NewXorm().Id(id).Update("is_delete", constant.Yes)
}