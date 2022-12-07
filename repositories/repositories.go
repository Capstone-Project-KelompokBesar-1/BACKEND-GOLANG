package repositories

import "ourgym/models"

type UserRepository interface {
	GetAll() []models.User
	GetOneByFilter(key string, value any) models.User
	Create(userRequest models.User) models.User
	Update(id string, userRequest models.User) models.User
	Delete(id uint) bool
}

type OtpRepository interface {
	GetOneByFilter(key string, value any) models.Otp
}
