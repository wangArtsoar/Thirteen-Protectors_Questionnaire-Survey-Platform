package facade

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"

type IServerService interface {
	SaveServer(server *models.Server, email string) error
	FindAllServerByUserEmail(userEmail string) ([]*models.Server, error)
	FindAllChannelByServerId(serverId int64) ([]*models.Channel, error)
	SaveChannel(channel *models.Channel) error
	SaveServerMember(serverMember *models.ServerMember) error
	SaveIdentity(identity *models.Identity) error
	SaveMemberRole(role *models.MemberRole) error
}
