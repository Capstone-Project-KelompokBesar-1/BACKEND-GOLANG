package services

import (
	"errors"
	"math/rand"
	"net/smtp"
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

func (as *AuthServiceImpl) ForgotPassword(email string) error {
	user := as.userRepo.GetOneByFilter("email", email)

	if user.Email == "" {
		return errors.New("user not found")
	}

	rand.Seed(time.Now().UnixNano())
	randCode := rand.Intn(9999-0000) + 0000
	otpRequest := models.Otp{
		UserID:    user.ID,
		Code:      randCode,
		ExpiredAt: time.Now().Add(time.Minute * 3),
	}

	otp := as.otpRepo.Create(otpRequest)

	auth := smtp.PlainAuth(
		"",
		"fadlieferdiansyah62@gmail.com",
		"ejcvjptcyesolcox",
		"smtp.gmail.com",
	)
	msg := "Subject: Ourgym: OTP forgot password\n" + "This is your otp code " + strconv.Itoa(otp.Code) + ", please take care of your code\n *the otp code will be expired in 3 minutes"
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		user.Email,
		[]string{user.Email},
		[]byte(msg),
	)

	if err != nil {
		return err
	}

	return nil
}

func (as *AuthServiceImpl) ValidateOTP(otpCode int) (map[string]string, error) {
	otp := as.otpRepo.GetOneByFilter("code", otpCode)

	if otp.ID == 0 {
		return nil, errors.New("otp invalid")
	}

	timeNow := time.Now().Unix()
	expiredAt := otp.ExpiredAt.Unix()

	if expiredAt < timeNow {
		return nil, errors.New("otp expired")
	}

	as.otpRepo.Delete(otp.ID)

	user := as.userRepo.GetOneByFilter("id", otp.UserID)
	token, _ := middlewares.GenerateToken(user, 1)

	return map[string]string{
		"token": token,
	}, nil
}

func (as *AuthServiceImpl) ResetPassword(id, new_password string) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)

	isSuccess := as.userRepo.ChangePassword(id, string(password))

	if !isSuccess {
		return errors.New("failed to create new password")
	}

	return nil
}
