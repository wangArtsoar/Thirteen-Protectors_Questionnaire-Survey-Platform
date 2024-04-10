package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/facade"
	"github.com/go-xorm/xorm"
)

var _ facade.IChannelRepo = new(ChannelRepo)

type ChannelRepo struct {
}

// NewChannelRepo 构造器
func NewChannelRepo() *ChannelRepo {
	return &ChannelRepo{}
}

func (c *ChannelRepo) CreateChannel(session *xorm.Session, channel *models.Channel) (int64, error) {
	return session.InsertOne(channel)
}

func (c *ChannelRepo) FindAllByServerId(limit int, serverId int64) ([]*models.Channel, error) {
	var channels []*models.Channel
	query := persistence.NewXorm().Where("server_id = ?", serverId)
	if limit != 0 {
		query = query.Limit(limit)
	}
	return channels, query.Find(&channels)
}

func (c *ChannelRepo) FindOneByChannelName(channelName string) (*models.Channel, error) {
	var channel models.Channel
	_, err := persistence.NewXorm().Where("channel_name = ?", channelName).Get(channel)
	return &channel, err
}
