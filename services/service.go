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
	ChangePassword(id string, newPassword string) bool
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type ClassService interface {
	GetAll(classType string, name string) []models.Class
	GetByID(id string) models.Class
	Create(classRequest models.Class) models.Class
	Update(id string, classRequest models.Class) models.Class
	Delete(id string) bool
	DeleteMany(ids string) bool
}
