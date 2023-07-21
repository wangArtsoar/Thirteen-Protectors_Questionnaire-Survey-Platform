package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/dao"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service/facade"
	"go.uber.org/dig"
)

var (
	c = dig.New()
	C Container
)

type Container struct {
	UserService   facade.IUserService   `dig:"UserService"`
	ServerService facade.IServerService `dig:"ServerService"`
}

func InitIoc() {
	initRepos()
	initService()
	initCtrl()
}

// 仓库依赖
func initRepos() {
	handleErr(c.Provide(dao.NewChannelRepo))
	handleErr(c.Provide(dao.NewIdentityRepo))
	handleErr(c.Provide(dao.NewLabelRepo))
	handleErr(c.Provide(dao.NewMemberRoleRepo))
	handleErr(c.Provide(dao.NewServerMemberRepo))
	handleErr(c.Provide(dao.NewMessgeRepo))
	handleErr(c.Provide(dao.NewServerRepo))
	handleErr(c.Provide(user.NewUserRepo))
}

// 服务依赖
func initService() {
	handleErr(c.Provide(service.NewServerService))
	handleErr(c.Provide(service.NewUserService))
}

// 容器接口
func initCtrl() {
	handleErr(c.Invoke(func(userService facade.IUserService) {
		C.UserService = userService
	}))
	handleErr(c.Invoke(func(serverService facade.IServerService) {
		C.ServerService = serverService
	}))
}

// 错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
