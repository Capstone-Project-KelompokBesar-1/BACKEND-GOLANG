package services

import (
	"errors"
	"ourgym/middlewares"
	"ourgym/models"
	"ourgym/repositories"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func (as *AuthServiceImpl) Login(userRequest models.User) (map[string]string, error) {
	user := as.userRepo.GetOneByFilter("email", userRequest.Email)

	if user.ID == 0 {
		return map[string]string{}, errors.New("email hasn't been registered")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		return map[string]string{}, errors.New("password invalid")
	}

	token, _ := middlewares.GenerateToken(user, 6)
	refreshToken, _ := middlewares.GenerateToken(user, 12)

	return map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	}, nil
}

func (as *AuthServiceImpl) Register(userRequest models.User) error {
	user := as.userRepo.GetOneByFilter("email", userRequest.Email)

	if userRequest.Email == user.Email {
		return errors.New("email has been registered")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)

	userRequest.Password = string(password)
	userRequest.IsAdmin = false

	user = as.userRepo.Create(userRequest)

	return nil
}
