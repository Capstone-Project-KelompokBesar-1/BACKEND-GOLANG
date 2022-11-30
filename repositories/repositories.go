package repositories

import "ourgym/models"

type UserRepository interface {
	GetAll(name string) []models.User
	GetOneByFilter(key string, value any) models.User
	Create(userRequest models.User) models.User
	Update(id string, userRequest models.User) models.User
	Delete(id string) bool
	DeleteMany(ids string) bool
}
