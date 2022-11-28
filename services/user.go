package services

import (
	"ourgym/models"
	"ourgym/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserServices {
	return &UserService{
		userRepository: ur,
	}
}

func (u *UserService) GetAll() []models.User {
	return u.userRepository.GetAll()
}

func (u *UserService) GetOneByFilter(key string, value any) models.User {
	return u.userRepository.GetOneByFilter(key, value)
}

func (u *UserService) Create(userRequest models.User) models.User {
	return u.userRepository.Create(userRequest)
}

func (u *UserService) Update(id string, userRequest models.User) models.User {
	return u.userRepository.Update(id, userRequest)
}

func (u *UserService) Delete(id string) bool {
	return u.userRepository.Delete(id)
}
