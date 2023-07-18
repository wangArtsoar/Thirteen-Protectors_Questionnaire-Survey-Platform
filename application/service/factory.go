package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"time"
)

func ToMemberRoleModel(serverID int64) *models.MemberRole {
	return &models.MemberRole{
		ServerId:    serverID,
		Name:        constant.OWNER,
		Permissions: constant.Owner(),
	}
}

func ToIdentityModel(serverID int64) *models.Identity {
	return &models.Identity{
		ServerId:   serverID,
		Name:       constant.Default,
		MemberRole: constant.OWNER,
	}
}

func ToServerModel(id, email, name string) *models.Server {
	if len(name) < 1 {
		name = constant.Default
	}
	return &models.Server{
		Name:       name,
		CreateAt:   time.Now(),
		OwnerId:    id,
		OwnerEmail: email,
	}
}

func ToLabelModel(serverId int64, label string) *models.Label {
	return &models.Label{
		ServerId: serverId,
		Name:     label,
	}
}

func ToServerMemberModel(serverId int64, identityID int64, userInfo *models.User) *models.ServerMember {
	return &models.ServerMember{
		ServerId:   serverId,
		MemberName: userInfo.Name,
		UserName:   userInfo.Name,
		UserId:     userInfo.ID,
		UserEmail:  userInfo.Email,
		CreateAt:   time.Now(),
		IdentityId: identityID,
	}
}
