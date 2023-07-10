package facade

import "Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"

type IIdentityRepo interface {
	SaveIdentity(identity *models.Identity) (int64, error)
	FindAllIdentityByChannelId(channelId int64, err error) ([]*models.Identity, error)
}
