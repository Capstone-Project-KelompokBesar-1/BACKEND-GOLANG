package services

import (
	"ourgym/models"
)

type AuthService interface {
	Login(userRequest models.User) (map[string]string, error)
	Register(userRequest models.User) error
}
