package dao

import "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/facade"

var _ facade.ILabelRepo = new(LabelRepo)

type LabelRepo struct {
}
