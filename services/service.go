package services

import (
	"ourgym/models"
)

type AuthService interface {
	Login(userRequest models.User) (map[string]string, error)
	Register(userRequest models.User) error
}

type UserServices interface {
	GetAll() []models.User
	GetOneByFilter(key string, value any) models.User
	Create(userRequest models.User) models.User
	Update(id string, userRequest models.User) models.User
	Delete(id string) bool
}
