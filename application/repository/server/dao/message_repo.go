package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
)

var _ facade.IMessageRepo = new(MessageRepo)

type MessageRepo struct {
}

func (m *MessageRepo) UpdateMessage(id int64, message *models.Message) (int64, error) {
	return orm.NewXorm().Id(id).Update(message)
}

func (m *MessageRepo) FindLastRecords(limit int) ([]*models.Message, error) {
	var messages []*models.Message
	return messages,
		orm.NewXorm().
			Where("is_withdraw = ?", constant.Default).
			Asc("date").
			Limit(limit).
			Find(messages)
}

func (m *MessageRepo) FindByKeywords(keywords string) ([]*models.Message, error) {
	var messages []*models.Message
	return messages, orm.NewXorm().Where("message LIKE ?", "%"+keywords+"%").Find(messages)
}
