package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	pl "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
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
	FindJoinServerListByUser(userEmail string, request pl.PageRequest) (*pl.PageList[models.Server], error)
}
