package ass

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/vo"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	pl "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
	"github.com/goccy/go-json"
	"github.com/samber/lo"
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

// ServerModelToServerResponse 将 ServerModel 转换成 ServerResponse
func ServerModelToServerResponse(serverModel models.Server) vo.ServerResponse {
	var labels []string
	json.Unmarshal(serverModel.Labels, labels)
	return vo.ServerResponse{
		ID:     serverModel.Id,
		Name:   serverModel.Name,
		Labels: labels,
	}
}

// PageServerModelToServerResponse 将 page 从 model 类型转换成 vo
func PageServerModelToServerResponse(list *pl.PageList[models.Server]) *pl.PageList[vo.ServerResponse] {
	return &pl.PageList[vo.ServerResponse]{
		PageNum:   list.PageNum,
		PageSize:  list.PageSize,
		PageTotal: list.PageTotal,
		Body: lo.Map(list.Body, func(item models.Server, index int) vo.ServerResponse {
			return ServerModelToServerResponse(item)
		}),
	}
}
