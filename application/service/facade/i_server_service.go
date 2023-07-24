package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/page_list"
)

type IServerService interface {
	SaveServer(server *models.Server, email string) error
	FindAllServerByUserEmail(userEmail string) ([]*models.Server, error)
	FindAllChannelByServerId(serverId int64) ([]*models.Channel, error)
	SaveChannel(channel *models.Channel) error
	SaveServerMember(serverMember *models.ServerMember) error
	SaveIdentity(identity *models.Identity) error
	SaveMemberRole(role *models.MemberRole) error
	SaveMessage(message *models.Message, userEmail string) (*models.Message, error)
	FindMessageByKeyword(keyword string) ([]*models.Message, error)
	FindMessageByLimit(limit int) ([]*models.Message, error)
	FindJoinServerListByUser(userEmail string, request page_list.PageRequest) (page_list.PageList[[]models.Server], error)
}
