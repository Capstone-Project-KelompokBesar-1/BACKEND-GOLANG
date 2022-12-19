package services

import (
	"errors"
	"ourgym/dto"
	"ourgym/models"
	"ourgym/repositories"

	"golang.org/x/crypto/bcrypt"
)

func NewUserService(ur repositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: ur,
	}
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func (u *UserServiceImpl) GetAll(name string) []dto.UserResponse {
	users := u.userRepository.GetAll(name)

	var usersResponse []dto.UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, user.ConvertToDTO())
	}

	return usersResponse
}

func (u *UserServiceImpl) GetByID(id string) dto.UserResponse {
	user := u.userRepository.GetOneByFilter("id", id)
	return user.ConvertToDTO()
}

func (u *UserServiceImpl) Create(userRequest dto.UserRequest) (dto.UserResponse, error) {
	userChecked := u.userRepository.GetOneByFilter("email", userRequest.Email)

	if userRequest.Email == userChecked.Email {
		return dto.UserResponse{}, errors.New("email has been registered")
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	userRequest.Password = string(newPassword)

	userModel := models.FromUserRequestToUserModel(userRequest)

	user := u.userRepository.Create(userModel)
	return user.ConvertToDTO(), nil
}

func (u *UserServiceImpl) Update(id string, userRequest dto.UserRequest) dto.UserResponse {
	userModel := models.FromUserRequestToUserModel(userRequest)

	user := u.userRepository.Update(id, userModel)

	return user.ConvertToDTO()
}

func (u *UserServiceImpl) ChangePassword(id string, passwords dto.ChangePasswordRequest) error {
	user := u.userRepository.GetOneByFilter("id", id)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwords.OldPassword))

	if err != nil {
		return errors.New("old password invalid")
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwords.NewPassword), bcrypt.DefaultCost)

	isSuccess := u.userRepository.ChangePassword(id, string(newPassword))

	if !isSuccess {
		return errors.New("failed to change password")
	}

	return nil
}

func (u *UserServiceImpl) UpdatePhoto(id string, userRequest dto.UserRequest) dto.UserResponse {
	userModel := models.FromUserRequestToUserModel(userRequest)

	user := u.userRepository.Update(id, userModel)

	return user.ConvertToDTO()
}

func (u *UserServiceImpl) Delete(id string) bool {
	return u.userRepository.Delete(id)
}

func (u *UserServiceImpl) DeleteMany(ids string) bool {
	return u.userRepository.DeleteMany(ids)
}
