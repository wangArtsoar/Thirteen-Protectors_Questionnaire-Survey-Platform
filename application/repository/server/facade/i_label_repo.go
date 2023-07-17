package facade

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"github.com/go-xorm/xorm"
)

type ILabelRepo interface {
	SaveLabel(session *xorm.Session, label []*models.Label) (int64, error)
	FindAll() ([]*models.Label, error)
	ExistByName(name string) (bool, error)
}
