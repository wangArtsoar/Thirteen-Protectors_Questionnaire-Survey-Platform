package user

import "Thirteen-Protectors_Questionnaire-Survey-Platform/models"

type IUserRepo interface {
	ExistByEmail(email string) (bool, error)
	FindByEmail(email string) (*models.User, error)
	SaveUser(user models.User) (int64, error)
}
