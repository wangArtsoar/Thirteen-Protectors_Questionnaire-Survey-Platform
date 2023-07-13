package ass

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
)

// MessageModelToResponse 将 MessageModel 转换成 Response
func MessageModelToResponse(message *models.Message) vo.MessageResponse {
	return vo.MessageResponse{
		Content:  message.Content,
		SendDate: message.SendDate,
	}
}

// MessageModelToResponseList 将 MessageModelList 转换成 ResponseList
func MessageModelToResponseList(messageModelList []*models.Message) []*vo.MessageResponse {
	var responseList []*vo.MessageResponse
	for _, message := range messageModelList {
		response := MessageModelToResponse(message)
		responseList = append(responseList, &response)
	}
	return responseList
}
