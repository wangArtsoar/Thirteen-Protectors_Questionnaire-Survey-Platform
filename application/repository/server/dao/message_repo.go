package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"fmt"
	"xorm.io/builder"
)

var _ facade.IMessageRepo = new(MessageRepo)

type MessageRepo struct {
}

func (m *MessageRepo) SaveMessage(message *models.Message) (*models.Message, error) {
	_, err := orm.NewXorm().InsertOne(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (m *MessageRepo) UpdateMessage(id int64, message *models.Message) (int64, error) {
	return orm.NewXorm().Id(id).Update(message)
}

func (m *MessageRepo) FindLastRecords(limit int) ([]*models.Message, error) {
	var messages []*models.Message
	sql, args, err := builder.Select().From("message").
		Where(builder.Eq{"is_withdraw": constant.No}).
		OrderBy("send_date ASC").ToSQL()
	if err != nil {
		return nil, err
	}
	limitInSql := limit
	if limit == 0 {
		limitInSql = 10
	}
	// 修改 sql 变量，添加 LIMIT 和 OFFSET 子句
	sql = fmt.Sprintf("%s LIMIT %d ", sql, limitInSql)
	return messages, orm.NewXorm().Find(messages, sql, args)
}

func (m *MessageRepo) FindByKeywords(keywords string) ([]*models.Message, error) {
	var messages []*models.Message
	return messages, orm.NewXorm().Where("type = ? and content LIKE ?", constant.No, "%"+keywords+"%").
		Find(messages)
}
