package ass

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
)

// ServerRequestToModel 请求体转换model
func ServerRequestToModel(request vo.ServerRequest) *models.Server {
	return &models.Server{
		Name:   request.Name,
		Labels: request.Labels,
	}
}
