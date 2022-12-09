package controllers

import (
	"fmt"
	"net/http"
	"ourgym/models"
	"ourgym/services"

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
	var userRequest models.User

	c.Bind(&userRequest)

	tokens, err := ac.authService.Login(userRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprint(err),
		})
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success logged in", tokens))
}

func (ac *AuthController) Register(c echo.Context) error {
	var userRequest models.User

	c.Bind(&userRequest)

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
	fmt.Println("Ini con : ", c.Get("user"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprint(err),
		})
	}
	return c.JSON(http.StatusOK, Response(http.StatusOK, "Password has been changed", map[string]any{}))
}
