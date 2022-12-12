package services

import (
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"ourgym/databases"
	"ourgym/dto"
	"ourgym/middlewares"
	"ourgym/models"
	"ourgym/repositories"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(userRepo repositories.UserRepository, otpRepo repositories.OtpRepository) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
	otpRepo  repositories.OtpRepository
}

func (as *AuthServiceImpl) Login(loginRequest dto.LoginRequest) (map[string]string, error) {
	user := as.userRepo.GetOneByFilter("email", loginRequest.Email)

	if user.ID == 0 {
		return map[string]string{}, errors.New("email hasn't been registered")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

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

func (as *AuthServiceImpl) Register(userRequest dto.UserRequest) error {
	user := as.userRepo.GetOneByFilter("email", userRequest.Email)

	if userRequest.Email == user.Email {
		return errors.New("email has been registered")
	}

	userModel := models.FromUserRequestToUserModel(userRequest)

	password, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)

	userModel.Password = string(password)

	userModel.IsAdmin = false

	user = as.userRepo.Create(userModel)

	return nil
}

func (as *AuthServiceImpl) SendOTP(email string) error {
	// var user models.User
	user := as.userRepo.GetOneByFilter("email", email)
	if user.Email != "" {
		token, _ := middlewares.GenerateToken(user, 1)
		rand.Seed(time.Now().UnixNano())
		randCode := rand.Intn(9999-0000) + 0000
		otp := models.Otp{
			User:  user.ID,
			Code:  strconv.Itoa(randCode),
			Token: token,
		}
		databases.InitDatabase().Create(&otp)

		auth := smtp.PlainAuth(
			"",
			"fadlieferdiansyah62@gmail.com",
			"ejcvjptcyesolcox",
			"smtp.gmail.com",
		)
		msg := "Subject: Ourgym: OTP forgot password\n" + "This is your otp code " + strconv.Itoa(randCode) + ", please take care of your code"
		err := smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			"fadlieferdiansyah62@gmail.com",
			[]string{user.Email},
			[]byte(msg),
		)

		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func (as *AuthServiceImpl) CreateNewPassword(otp, new_password string) error {

	oo := as.otpRepo.GetOneByFilter("code", otp)
	user := as.userRepo.GetOneByFilter("id", oo.User)
	id := strconv.Itoa(int(user.ID))
	password, _ := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	user.Password = string(password)
	as.userRepo.Update(id, user)
	as.otpRepo.Delete(oo.ID)

	return nil
}
