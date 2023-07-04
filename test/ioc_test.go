package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/service"
	"fmt"
	"github.com/facebookgo/inject"
	"testing"
)

func TestIoc(t *testing.T) {

	var container2 struct {
		UserService facade.IUserService `inject:"UserService"`
	}

	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &container2},
		&inject.Object{Value: &service.UserService{}, Name: "UserService"},
		&inject.Object{Value: &repository.UserRepo{}, Name: "UserRepo"},
	)
	if err != nil {
		panic(err)
	}
	err = g.Populate()
	if err != nil {
		panic(err)
	}

	userService := container2.UserService
	if userService == nil {
		fmt.Println("获取不到")
	} else {
		fmt.Println("获取到了")
	}
}
