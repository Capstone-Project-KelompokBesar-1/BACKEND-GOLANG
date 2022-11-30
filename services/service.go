package services

import (
	"ourgym/models"
)

type AuthService interface {
	Login(userRequest models.User) (map[string]string, error)
	Register(userRequest models.User) error
}

type UserService interface {
	GetAll(name string) []models.User
	GetByID(id string) models.User
	Create(userRequest models.User) models.User
	Update(id string, userRequest models.User) models.User
	Delete(id string) bool
	DeleteMany(ids string) bool
}
