package orm

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	_ "database/sql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"time"
	"xorm.io/core"
)

var (
	engine *xorm.Engine
	err    error
)

func NewXorm() *xorm.Engine {
	connStr := "postgres://postgres:xiaoyi@localhost/questionnaire_survey?sslmode=disable"
	engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Fatalf("ping to db fail! err:%+v", err)
		return nil
	}
	// 连接池配置
	engine.SetMaxOpenConns(30)                  // 最大 db 连接
	engine.SetMaxIdleConns(10)                  // 最大 db 连接空闲数
	engine.SetConnMaxLifetime(30 * time.Minute) // 超过空闲数连接存活时间

	// 日志相关配置
	engine.ShowSQL(true)                     // 打印日志
	engine.Logger().SetLevel(core.LOG_DEBUG) // 打印日志级别
	//engine.SetLogger(nil)                    // 设置日志输出 (控制台, 日志文件, 系统日志等)
	// 测试连通性
	if err = engine.Ping(); err != nil {
		log.Fatalf("ping to db fail! err:%+v", err)
		return nil
	}
	// 初始化一个管理角色
	var count int64
	if count, err = engine.Count(&models.Role{}); err != nil {
		log.Println("初始化角色失败")
		return nil
	}
	if count <= 0 {
		_, err = engine.InsertOne(&models.Role{Name: "ADMIN"})
		if err != nil {
			log.Println("初始化角色失败" + err.Error())
			return nil
		}
	}
	return engine
}
