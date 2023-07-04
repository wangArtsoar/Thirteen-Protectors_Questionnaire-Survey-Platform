package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/service"
	"github.com/facebookgo/inject"
)

func InitIndexIoc(g inject.Graph) {
	handleErr(g.Provide(
		&inject.Object{Value: &Container},
		&inject.Object{Value: &service.UserService{}, Name: "UserService"},
		&inject.Object{Value: &repository.UserRepo{}, Name: "UserRepo"},
	))
}
