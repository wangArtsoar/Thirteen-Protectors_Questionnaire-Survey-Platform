package ioc

import (
	repo "Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	srv "Thirteen-Protectors_Questionnaire-Survey-Platform/service"
	"github.com/facebookgo/inject"
)

func InitIndexIoc(g *inject.Graph) {
	handleErr(g.Provide(
		&inject.Object{Value: &srv.UserService{}, Name: "UserService"},
		&inject.Object{Value: &repo.UserRepo{}, Name: "UserRepo"},
	))
}
