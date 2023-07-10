package init

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"time"
)

func Server(name string) models.Server {
	return models.Server{
		Name:     name + "'s Server",
		CreateAt: time.Now(),
		Labels:   []string{"default"},
	}
}
