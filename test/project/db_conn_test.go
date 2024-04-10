package project

import (
	_ "database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

func TestDBConnect(t *testing.T) {
	connStr := "postgres://postgres:xiaoyi@localhost/questionnaire_survey"
	engine, err := xorm.NewEngine("postgres", connStr)
	if err != nil {
		fmt.Println("连接失败")
		log.Fatalf("ping to db fail! err:%+v", err)
	}
	if engine == nil {
		fmt.Println("空的")
	}
	fmt.Println("连接成功")
}
