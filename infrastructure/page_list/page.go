package page_list

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"github.com/goccy/go-json"
)

type PageRequest struct {
	PageNum     int `json:"page_num"`
	PageSize    int `json:"page_size"`
	CurrenTotal int `json:"curren_total"`
}

type PageList[T any] struct {
	PageNum   int `json:"page_num"`
	PageSize  int `json:"page_size"`
	PageTotal int `json:"page_total"`
	Body      []T `json:"body"`
}

func Pageable[T any](M []T, request PageRequest, count int) PageList[T] {
	return PageList[T]{
		PageNum:   request.PageNum,
		PageSize:  request.PageSize,
		PageTotal: count,
		Body:      M,
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
		PageNum:     num,
		PageSize:    size,
		CurrenTotal: num * size,
	}
}
