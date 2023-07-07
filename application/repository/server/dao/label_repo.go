package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
)

var _ facade.ILabelRepo = new(LabelRepo)

type LabelRepo struct {
}
