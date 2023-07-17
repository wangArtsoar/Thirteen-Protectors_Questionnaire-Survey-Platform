package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/router"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
)

func main() {
	r := router.Router()
	ioc.InitIoc()
	_ = r.Run(":8080")
}

//func init() {
//	fmt.Println("database init start")
//	if err := orm.NewXorm().Sync2(
//		models.User{},
//		models.Server{},
//		models.Channel{},
//		models.ServerMember{},
//		models.Label{},
//		models.Message{},
//		models.Identity{},
//		models.MemberRole{},
//	); err != nil {
//		_ = fmt.Errorf("database error : %v", err)
//	}
//}
