package main

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/page_list"
	"fmt"
	"testing"
)

func TestPage(T *testing.T) {
	page := page_list.DefaultPage("1", "3")
	// 获取label
	var labels []models.Label
	err := orm.NewXorm().Table("label").Limit(page.PageSize, page.PageNum-1).Find(&labels)
	if err != nil {
		T.Error(err)
	}
	pageable := page_list.Pageable(labels, page, 10)
	fmt.Println(pageable)
}
