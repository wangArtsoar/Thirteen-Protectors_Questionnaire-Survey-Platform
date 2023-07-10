package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
)

var _ facade.IChannelRepo = new(ChannelRepo)

type ChannelRepo struct {
}

func (c *ChannelRepo) CreateChannel(channel *models.Channel) (int64, error) {
	return orm.NewXorm().InsertOne(channel)
}

func (c *ChannelRepo) FindAllByServerId(serverId int64) ([]*models.Channel, error) {
	var channels []*models.Channel
	return channels, orm.NewXorm().Where("server_id = ?", serverId).Find(channels)
}

func (c *ChannelRepo) FindOneByChannelName(channelName string) (*models.Channel, error) {
	var channel models.Channel
	_, err := orm.NewXorm().Where("channel_name = ?", channelName).Get(channel)
	return &channel, err
}
