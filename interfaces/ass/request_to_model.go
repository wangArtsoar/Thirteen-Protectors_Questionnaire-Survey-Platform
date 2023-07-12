package ass

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
)

// ServerRequestToModel 将ServerRequest请求体转换model
func ServerRequestToModel(request vo.ServerRequest) *models.Server {
	return &models.Server{
		Name:   request.Name,
		Labels: request.Labels,
	}
}

// IdentityRequestToModel 将IdentityRequest请求体转换成model
func IdentityRequestToModel(request vo.IdentityRequest) *models.Identity {
	return &models.Identity{
		Name:       request.Name,
		ChannelId:  request.ChannelID,
		MemberRole: request.MemberRole,
	}
}

// MemberRoleRequestToModel 将MemberRoleRequest请求转换成model
func MemberRoleRequestToModel(request vo.MemberRoleRequest) *models.MemberRole {
	return &models.MemberRole{
		Name:        request.Name,
		ChannelId:   request.ChannelID,
		Permissions: request.Permissions,
	}
}
