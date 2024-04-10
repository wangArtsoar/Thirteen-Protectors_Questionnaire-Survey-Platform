package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/domain/server/service"
	service2 "Thirteen-Protectors_Questionnaire-Survey-Platform/domain/user/service"
	dao2 "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/dao"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/user"
	"go.uber.org/dig"
)

var (
	c = dig.New()
	C Container
)

type Container struct {
	UserService   service2.IUserService  `dig:"UserService"`
	ServerService service.IServerService `dig:"ServerService"`
}

func InitIoc() {
	initRepos()
	initService()
	initCtrl()
}

// 仓库依赖
func initRepos() {
	handleErr(c.Provide(dao2.NewChannelRepo))
	handleErr(c.Provide(dao2.NewIdentityRepo))
	handleErr(c.Provide(dao2.NewLabelRepo))
	handleErr(c.Provide(dao2.NewMemberRoleRepo))
	handleErr(c.Provide(dao2.NewServerMemberRepo))
	handleErr(c.Provide(dao2.NewMessgeRepo))
	handleErr(c.Provide(dao2.NewServerRepo))
	handleErr(c.Provide(user.NewUserRepo))
}

// 服务依赖
func initService() {
	handleErr(c.Provide(service.NewServerService))
	handleErr(c.Provide(service2.NewUserService))
}

// 容器接口
func initCtrl() {
	handleErr(c.Invoke(func(userService service2.IUserService) {
		C.UserService = userService
	}))
	handleErr(c.Invoke(func(serverService service.IServerService) {
		C.ServerService = serverService
	}))
}

// 错误处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
