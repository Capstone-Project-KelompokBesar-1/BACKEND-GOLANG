package repositories

import "ourgym/models"

type UserRepository interface {
	GetAll(name string) []models.User
	GetOneByFilter(key string, value any) models.User
	Create(userRequest models.User) models.User
	Update(id string, userRequest models.User) models.User
	ChangePassword(id string, newPassword string) bool
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type ClassRepository interface {
	GetAll(classType string, name string) []models.Class
	GetOneByFilter(key string, value any) models.Class
	Create(classRequest models.Class) models.Class
	Update(id string, userRequest models.Class) models.Class
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type TrainerRepository interface {
	GetAll(name string) []models.Trainer
	GetByID(id string) models.Trainer
}

type OtpRepository interface {
	GetOneByFilter(key string, value any) models.Otp
	Delete(id uint) bool
}
