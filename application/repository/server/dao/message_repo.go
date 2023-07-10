package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
)

var _ facade.IMessageRepo = new(MessageRepo)

type MessageRepo struct {
}

func (m *MessageRepo) UpdateMessage(id int64, message *models.Message) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageRepo) FindLastRecords(limit int) ([]*models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MessageRepo) FindByKeywords(keywords string) ([]*models.Message, error) {
	//TODO implement me
	panic("implement me")
}
