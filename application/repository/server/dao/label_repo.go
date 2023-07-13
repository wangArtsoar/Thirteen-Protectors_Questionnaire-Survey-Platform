package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/go-xorm/xorm"
)

var _ facade.ILabelRepo = new(LabelRepo)

type LabelRepo struct {
}

func (l *LabelRepo) ExistByName(session *xorm.Session, name string) (bool, error) {
	return session.Where("name = ?", name).Exist()
}

func (l *LabelRepo) SaveLabel(label []*models.Label) (int64, error) {
	return orm.NewXorm().Insert(label)
}

func (l *LabelRepo) FindAll() ([]*models.Label, error) {
	var labels []*models.Label
	return labels, orm.NewXorm().Find(labels)
}
