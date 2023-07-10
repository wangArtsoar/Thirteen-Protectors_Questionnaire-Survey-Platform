package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
)

var _ facade.IServerMemberRepo = new(ServerMemberRepo)

type ServerMemberRepo struct {
}

func (s *ServerMemberRepo) NewAServerMember(member *models.ServerMember) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerMemberRepo) FindAllServerMember() ([]*models.ServerMember, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerMemberRepo) FindAllServerMemberByServerId(serverId int64) ([]*models.ServerMember, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerMemberRepo) FindAllServerMemberByChannel(channel *models.Channel) ([]*models.ServerMember, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerMemberRepo) FindServerMemberByIdentityId(identityId int64) ([]*models.ServerMember, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerMemberRepo) UpdateServerMemberByUserEmail(userEmail string, member *models.ServerMember) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerMemberRepo) DeleteServerMemberById(id int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}
