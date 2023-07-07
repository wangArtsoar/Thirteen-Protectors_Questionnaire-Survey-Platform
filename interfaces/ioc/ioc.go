package ioc

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/dao"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	service2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/service"
	facade2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/service/facade"
	"github.com/facebookgo/inject"
)

var Container struct {
	UserService   facade2.IUserService   `inject:"UserService"`
	ServerService facade2.IServerService `inject:"ServerService"`
}

func InitIoc() {
	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &Container},
		&inject.Object{Value: &service2.UserService{}, Name: "UserService"},
		&inject.Object{Value: &user.UserRepo{}, Name: "UserRepo"},

		&inject.Object{Value: &service2.ServerService{}, Name: "ServerService"},
		&inject.Object{Value: &dao.ServerRepo{}, Name: "ServerRepo"},
		&inject.Object{Value: &dao.ChannelRepo{}, Name: "ChannelRepo"},
		&inject.Object{Value: &dao.ServerMemberRepo{}, Name: "ServerMemberRepo"},
		&inject.Object{Value: &dao.LabelRepo{}, Name: "LabelRepo"},
		&inject.Object{Value: &dao.MemberRoleRepo{}, Name: "MemberRoleRepo"},
		&inject.Object{Value: &dao.IdentityRepo{}, Name: "IdentityRepo"},
		&inject.Object{Value: &dao.MessageRepo{}, Name: "MessageRepo"},
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
