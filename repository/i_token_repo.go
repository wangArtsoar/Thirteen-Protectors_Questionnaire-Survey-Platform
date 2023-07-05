package repository

import "Thirteen-Protectors_Questionnaire-Survey-Platform/models"

type ITokenRepo interface {
	SaveToken(token models.Token) (int64, error)
	ExistToken(jwtToken string) (bool, error)
	DeleteTokenByUserId(userId string) (int64, error)
}
