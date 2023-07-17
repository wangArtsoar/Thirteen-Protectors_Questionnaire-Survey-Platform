package ass

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
	"time"
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

// MessageRequestToModel 将 MessageRequest 转换成 model
func MessageRequestToModel(request vo.MessageRequest) *models.Message {
	return &models.Message{
		Type:        request.Type,
		Content:     request.Content,
		SendDate:    time.Now(),
		LimitedTime: time.Now().Add(time.Minute * 2),
	}
}

// ChannelRequestToModel 将 ChannelRequest 转换成 model
func ChannelRequestToModel(request vo.ChannelRequest, serverID int64) *models.Channel {
	return &models.Channel{
		Name:     request.Name,
		ServerId: serverID,
		Label:    request.Label,
		CreateAt: time.Now(),
	}
}

// ServerMemberRequestToModel 将 ServerMemberRequest 转换成 model
func ServerMemberRequestToModel(request vo.ServerMemberRequest, email string) *models.ServerMember {
	return &models.ServerMember{
		ServerId:   request.ServerID,
		MemberName: request.MemberName,
		InviteId:   request.InviteId,
		UserEmail:  email,
	}
}
