package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/go-xorm/xorm"
)

var _ facade.IIdentityRepo = new(IdentityRepo)

type IdentityRepo struct {
}

func (i *IdentityRepo) SaveIdentity(session *xorm.Session, identity *models.Identity) (int64, error) {
	if _, err := session.InsertOne(identity); err != nil {
		return 0, err
	}
	if _, err := session.SQL(`select last_value from identity_id_seq `).Get(&identity.Id); err != nil {
		return 0, err
	}
	return identity.Id, nil
}

func (i *IdentityRepo) FindAllIdentityByChannelId(channelId int64, err error) ([]*models.Identity, error) {
	var identities []*models.Identity
	return identities, orm.NewXorm().Where("channel_id = ?", channelId).Find(identities)
}
