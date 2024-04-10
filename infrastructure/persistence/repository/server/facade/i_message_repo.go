package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
)

type IMessageRepo interface {
	FindLastRecords(limit int) ([]*models.Message, error)
	FindByKeywords(keywords string) ([]*models.Message, error)
	UpdateMessage(id int64, message *models.Message) (int64, error)
	SaveMessage(message *models.Message) (*models.Message, error)
}
