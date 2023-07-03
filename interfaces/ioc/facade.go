package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/service"
	"github.com/facebookgo/inject"
)

var DIContainer Container

type Container struct {
	UserService service.IUserService `inject:"UserService"`
}

func InitIoc() {
	var g inject.Graph
	handleErr(g.Provide(&inject.Object{Value: &DIContainer}))
	InitIndexIoc(&g)
	handleErr(g.Populate())
}

// 错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
