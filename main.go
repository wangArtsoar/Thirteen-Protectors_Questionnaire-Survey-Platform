package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/router"
)

func main() {
	r := router.Router()
	_ = r.Run(":8080")
	ioc.InitIoc()
}
