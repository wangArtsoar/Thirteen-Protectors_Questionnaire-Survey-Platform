package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"github.com/go-xorm/xorm"
)

type ILabelRepo interface {
	SaveLabel(label []*models.Label) (int64, error)
	FindAll() ([]*models.Label, error)
	ExistByName(session *xorm.Session, name string) (bool, error)
}
