package page_list

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"github.com/goccy/go-json"
)

type PageRequest struct {
	PageNum  int
	PageSize int
}

type PageList[Body any] struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
	Body     any `json:"body"`
}

func Pageable[T any](M T, request PageRequest) PageList[T] {
	return PageList[T]{
		PageNum:  request.PageNum,
		PageSize: request.PageSize,
		Body:     M,
	}
}

func DefaultPage(numStr, sizeStr string) PageRequest {
	var num, size int
	if len(numStr) == 0 {
		num = constant.PageNum
	} else {
		json.Unmarshal([]byte(numStr), &num)
	}
	if len(sizeStr) == 0 {
		size = constant.PageSize
	} else {
		json.Unmarshal([]byte(sizeStr), &size)
	}
	return PageRequest{
		PageNum:  num,
		PageSize: size,
	}
}
