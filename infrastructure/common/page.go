package common

import "Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"

type PageRequest struct {
	PageNum  int
	PageSize int
}

type PageList[Body []any] struct {
	PageNum  int   `json:"page_num"`
	PageSize int   `json:"page_size"`
	Body     []any `json:"body"`
}

var P PageList[[]any]

func Pageable(M []any) PageList[[]any] {
	P = Page(constant.Zero, constant.Zero)
	P.Body = M
	return P
}

func Page(pageNum, pageSize int) PageList[[]any] {
	if pageNum == constant.Zero {
		pageNum = constant.PageNum
	}
	if pageSize == constant.Zero {
		pageSize = constant.PageSize
	}
	return PageList[[]any]{
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}
