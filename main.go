package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	model "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/router"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
	"fmt"
)

func main() {
	r := router.Router()
	ioc.InitIoc()
	_ = r.Run(":8080")
}

func init() {
	fmt.Println("database init start")
	if err := persistence.NewXorm().Sync2(
		model.User{},
		model.Server{},
		model.Channel{},
		model.ServerMember{},
		model.Label{},
		model.Message{},
		model.Identity{},
		model.MemberRole{},
	); err != nil {
		_ = fmt.Errorf("database error : %v", err)
	}
}
