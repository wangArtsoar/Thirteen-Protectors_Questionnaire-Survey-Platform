package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/facade"
	"github.com/go-xorm/xorm"
)

var _ facade.ILabelRepo = new(LabelRepo)

type LabelRepo struct {
}

func NewLabelRepo() *LabelRepo {
	return &LabelRepo{}
}

func (l *LabelRepo) ExistByName(name string) (bool, error) {
	return persistence.NewXorm().
		Table(constant.LabelTable).
		Where("name = ?", name).Exist()
}

func (l *LabelRepo) SaveLabel(session *xorm.Session, label []*models.Label) (int64, error) {
	return session.Insert(label)
}

func (l *LabelRepo) FindAll() ([]*models.Label, error) {
	var labels []*models.Label
	return labels, persistence.NewXorm().Find(labels)
}
