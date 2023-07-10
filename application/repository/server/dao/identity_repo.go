package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
)

var _ facade.IIdentityRepo = new(IdentityRepo)

type IdentityRepo struct {
}

func (i *IdentityRepo) SaveIdentity(identity *models.Identity) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (i *IdentityRepo) FindAllIdentityByChannelId(channelId int64, err error) ([]*models.Identity, error) {
	//TODO implement me
	panic("implement me")
}
