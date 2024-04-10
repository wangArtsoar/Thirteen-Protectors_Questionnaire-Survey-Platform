package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/facade"
	"github.com/go-xorm/xorm"
)

var _ facade.IIdentityRepo = new(IdentityRepo)

type IdentityRepo struct {
}

// NewIdentityRepo 构造器
func NewIdentityRepo() *IdentityRepo {
	return &IdentityRepo{}
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
	return identities, persistence.NewXorm().Where("channel_id = ?", channelId).Find(identities)
}
