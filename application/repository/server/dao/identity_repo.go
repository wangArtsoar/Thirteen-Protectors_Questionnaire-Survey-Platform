package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
)

var _ facade.IIdentityRepo = new(IdentityRepo)

type IdentityRepo struct {
}

func (i *IdentityRepo) SaveIdentity(identity *models.Identity) (int64, error) {
	return orm.NewXorm().InsertOne(identity)
}

func (i *IdentityRepo) FindAllIdentityByChannelId(channelId int64, err error) ([]*models.Identity, error) {
	var identities []*models.Identity
	return identities, orm.NewXorm().Where("channel_id = ?", channelId).Find(identities)
}
