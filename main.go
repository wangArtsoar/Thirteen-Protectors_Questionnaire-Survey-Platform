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
