package repository

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/orm"
)

var _ ITokenRepo = new(TokenRepo)
var engine = orm.NewXorm()

type TokenRepo struct {
}

func (t *TokenRepo) DeleteTokenByUserId(userId string) (int64, error) {
	var token = models.Token{IsValid: 1}
	return engine.Where("user_id = ?", userId).Update(token)
}

func (t *TokenRepo) SaveToken(token models.Token) (int64, error) {
	return engine.InsertOne(token)
}

func (t *TokenRepo) ExistToken(jwtToken string) (bool, error) {
	return engine.Where("access_token = ? and is_valid = ?", jwtToken, 0).Exist(&models.Token{})
}
