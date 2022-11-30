package services

import (
	"ourgym/models"
	"ourgym/repositories"
)

func NewUserService(ur repositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: ur,
	}
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func (u *UserServiceImpl) GetAll(name string) []models.User {
	return u.userRepository.GetAll(name)
}

func (u *UserServiceImpl) GetByID(id string) models.User {
	return u.userRepository.GetOneByFilter("id", id)
}

func (u *UserServiceImpl) Create(userRequest models.User) models.User {
	return u.userRepository.Create(userRequest)
}

func (u *UserServiceImpl) Update(id string, userRequest models.User) models.User {
	return u.userRepository.Update(id, userRequest)
}

func (u *UserServiceImpl) UpdatePhoto(id string, userRequest models.User) models.User {
	return u.userRepository.Update(id, userRequest)
}

func (u *UserServiceImpl) Delete(id string) bool {
	return u.userRepository.Delete(id)
}

func (u *UserServiceImpl) DeleteMany(ids string) bool {
	return u.userRepository.DeleteMany(ids)
}
