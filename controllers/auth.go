package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"ourgym/dto"
	"ourgym/middlewares"
	"ourgym/models"
	"ourgym/services"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

type AuthController struct {
	authService services.AuthService
}

func (ac *AuthController) Login(c echo.Context) error {
	var loginRequest dto.LoginRequest

	c.Bind(&loginRequest)

	if err := loginRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	tokens, err := ac.authService.Login(loginRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, fmt.Sprint(err), nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success logged in", tokens))
}

func (ac *AuthController) Register(c echo.Context) error {
	var userRequest dto.UserRequest

	c.Bind(&userRequest)

	if err := userRequest.Validate(); err != nil || userRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	err := ac.authService.Register(userRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, fmt.Sprint(err), nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success registered user", map[string]any{}))
}

func (ac *AuthController) ForgotPassword(c echo.Context) error {
	var request struct {
		Email string `json:"email" form:"email" validate:"required,email"`
	}

	if err := c.Bind(&request); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if _, err := mail.ParseAddress(request.Email); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	err := ac.authService.ForgotPassword(request.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, fmt.Sprint(err), nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "OTP has been sended to email, Please check your email", map[string]any{}))
}

func (ac *AuthController) ValidateOTP(c echo.Context) error {
	var request struct {
		OTPCode int `json:"otp_code" form:"otp_code"`
	}

	if err := c.Bind(&request); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	otpCodeString := strconv.FormatInt(int64(request.OTPCode), 10)

	if len(otpCodeString) != 4 {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	otpCode, _ := strconv.Atoi(otpCodeString)

	token, err := ac.authService.ValidateOTP(otpCode)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, fmt.Sprint(err), nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "OTP code valid", token))
}

func (ac *AuthController) ResetPassword(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)

	userID := getUserIdFromToken(userToken)

	var passwords dto.ChangePasswordRequest

	if err := c.Bind(&passwords); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	err := ac.authService.ResetPassword(userID, passwords.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprint(err),
		})
	}
	return c.JSON(http.StatusOK, Response(http.StatusOK, "Password has been changed", map[string]any{}))
}

func (ac *AuthController) RefreshToken(c echo.Context) error {

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	user := models.User{
		ID:      uint(claims["id"].(float64)),
		IsAdmin: claims["is_admin"].(bool),
	}

	token, _ := middlewares.GenerateToken(user, 6)
	refresh_token, _ := middlewares.GenerateToken(user, 12)

	tokens := map[string]any{
		"token":         token,
		"refresh_token": refresh_token,
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success refresh token", tokens))
}
