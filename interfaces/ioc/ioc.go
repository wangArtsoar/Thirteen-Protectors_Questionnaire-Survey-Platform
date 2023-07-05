package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/service"
	"github.com/facebookgo/inject"
)

var Container struct {
	UserService facade.IUserService `inject:"UserService"`
}

func InitIoc() {
	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &Container},
		&inject.Object{Value: &service.UserService{}, Name: "UserService"},
		&inject.Object{Value: &repository.UserRepo{}, Name: "UserRepo"},
		&inject.Object{Value: &repository.RoleRepo{}, Name: "RoleRepo"},
		&inject.Object{Value: &repository.TokenRepo{}, Name: "TokenRepo"},
	)
	handleErr(err)
	err = g.Populate()
	handleErr(err)
}

// 错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

//func InitIndexIoc(g inject.Graph) {
//	handleErr(g.Provide(
//		&inject.Object{Value: &Container},
//		&inject.Object{Value: &service.UserService{}, Name: "UserService"},
//		&inject.Object{Value: &repository.UserRepo{}, Name: "UserRepo"},
//	))
//}
