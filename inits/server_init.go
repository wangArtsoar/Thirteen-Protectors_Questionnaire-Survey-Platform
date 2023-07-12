package inits

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
)

func Server(name string) models.Server {
	return models.Server{
		Name:   name + "'s Server",
		Labels: []string{"default"},
	}
}
