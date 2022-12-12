package controllers

import (
	"fmt"
	"net/http"
	"ourgym/dto"
	"ourgym/middlewares"
	"ourgym/models"
	"ourgym/services"

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
		return c.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprint(err),
		})
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success registered user", map[string]any{}))
}

func (ac *AuthController) SendOTP(c echo.Context) error {
	err := ac.authService.SendOTP(c.FormValue("email"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprint(err),
		})
	}
	return c.JSON(http.StatusOK, Response(http.StatusOK, "OTP has been sended to email, Please check your email", map[string]any{}))
}

func (ac *AuthController) CreateNewPassword(c echo.Context) error {
	err := ac.authService.CreateNewPassword(c.FormValue("code_otp"), c.FormValue("new_password"))
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
