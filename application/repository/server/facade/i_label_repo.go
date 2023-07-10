package facade

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"

type ILabelRepo interface {
	SaveLabel(label []*models.Label) (int64, error)
	FindAll() ([]*models.Label, error)
}
