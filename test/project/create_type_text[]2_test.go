package project

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func TestTypeText2(t *testing.T) {
	engine, err := xorm.NewEngine("postgres",
		"user=postgres password=xiaoyi dbname=questionnaire_survey sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(MyStruct))
	if err != nil {
		log.Fatal(err)
	}

	jsonStr := `{"MySlice": ["a", "b", "c"]}`
	var myStruct MyStruct
	err = json.Unmarshal([]byte(jsonStr), &myStruct)
	if err != nil {
		log.Fatal(err)
	}

	_, err = engine.Insert(&myStruct)
	if err != nil {
		log.Fatal(err)
	}
}
