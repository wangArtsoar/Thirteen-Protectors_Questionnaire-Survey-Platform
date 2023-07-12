package main

import (
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type MyStruct struct {
	Id      int64
	MySlice []string `xorm:"text"`
}

func TestCreate(t *testing.T) {
	engine, err := xorm.NewEngine("postgres",
		"user=postgres password=xiaoyi dbname=questionnaire_survey sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(MyStruct))
	if err != nil {
		log.Fatal(err)
	}

	mySlice := []string{"a", "b", "c"}
	_, err = engine.Insert(&MyStruct{MySlice: mySlice})
	if err != nil {
		log.Fatal(err)
	}

	var myStructs []MyStruct
	err = engine.Find(&myStructs)
	if err != nil {
		log.Fatal(err)
	}

	for _, myStruct := range myStructs {
		fmt.Println(myStruct.MySlice)
	}
}
