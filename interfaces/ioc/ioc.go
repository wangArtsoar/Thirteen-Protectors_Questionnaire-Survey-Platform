package ioc

import (
	server "Thirteen-Protectors_Questionnaire-Survey-Platform/repository/server/dao"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/service"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/service/facade"
	"github.com/facebookgo/inject"
)

var Container struct {
	UserService   facade.IUserService   `inject:"UserService"`
	ServerService facade.IServerService `inject:"ServerService"`
}

func InitIoc() {
	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &Container},
		&inject.Object{Value: &service.UserService{}, Name: "UserService"},
		&inject.Object{Value: &user.UserRepo{}, Name: "UserRepo"},

		&inject.Object{Value: &service.ServerService{}, Name: "ServerService"},
		&inject.Object{Value: &server.ServerRepo{}, Name: "ServerRepo"},
		&inject.Object{Value: &server.ChannelRepo{}, Name: "ChannelRepo"},
		&inject.Object{Value: &server.ServerMemberRepo{}, Name: "ServerMemberRepo"},
		&inject.Object{Value: &server.LabelRepo{}, Name: "LabelRepo"},
		&inject.Object{Value: &server.MemberRoleRepo{}, Name: "MemberRoleRepo"},
		&inject.Object{Value: &server.MemberRoleRepo{}, Name: "MemberRoleRepo"},
		&inject.Object{Value: &server.MessageRepo{}, Name: "MessageRepo"},
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
