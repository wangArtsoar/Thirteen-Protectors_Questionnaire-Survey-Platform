package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/facade"
	"github.com/facebookgo/inject"
)

var Container struct {
	UserService facade.IUserService `inject:"UserService"`
}

func InitIoc() {
	var g inject.Graph
	InitIndexIoc(g)
	handleErr(g.Populate())
}

// 错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
