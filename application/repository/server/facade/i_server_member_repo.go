package facade

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"

type IServerMemberRepo interface {
	NewAServerMember(member *models.ServerMember) (int64, error)
	FindAllServerMember() ([]*models.ServerMember, error)
	FindAllServerMemberByServerId(serverId int64) ([]*models.ServerMember, error)
	FindAllServerMemberByChannel(channel *models.Channel) ([]*models.ServerMember, error)
	FindServerMemberByIdentityId(identityId int64) ([]*models.ServerMember, error)
	UpdateServerMemberByUserEmail(userEmail string, member *models.ServerMember) (int64, error)
	DeleteServerMemberById(id int64) (int64, error)
}
