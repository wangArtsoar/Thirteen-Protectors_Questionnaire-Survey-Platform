package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	models2 "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"time"
)

func ToMemberRoleModel(serverID int64) *models2.MemberRole {
	return &models2.MemberRole{
		ServerId:    serverID,
		Name:        constant.OWNER,
		Permissions: constant.Owner(),
	}
}

func ToIdentityModel(serverID int64) *models2.Identity {
	return &models2.Identity{
		ServerId:   serverID,
		Name:       constant.Default,
		MemberRole: constant.OWNER,
	}
}

func ToServerModel(id, email, name string) *models2.Server {
	if len(name) < 1 {
		name = constant.Default
	}
	return &models2.Server{
		Name:       name,
		CreateAt:   time.Now(),
		OwnerId:    id,
		OwnerEmail: email,
	}
}

func ToLabelModel(serverId int64, label string) *models2.Label {
	return &models2.Label{
		ServerId: serverId,
		Name:     label,
	}
}

func ToServerMemberModel(serverId int64, identityID int64, userInfo *models2.User) *models2.ServerMember {
	return &models2.ServerMember{
		ServerId:   serverId,
		MemberName: userInfo.Name,
		UserName:   userInfo.Name,
		UserId:     userInfo.ID,
		UserEmail:  userInfo.Email,
		CreateAt:   time.Now(),
		IdentityId: identityID,
	}
}
