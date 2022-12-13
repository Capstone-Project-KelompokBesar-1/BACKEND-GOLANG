package controllers

import (
	"fmt"
	"net/http"
	"ourgym/dto"
	"ourgym/helpers"
	"ourgym/services"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	userService services.UserService
}

func NewProfileController(userService services.UserService) *ProfileController {
	return &ProfileController{
		userService,
	}
}

func (pc *ProfileController) GetProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)

	userID := getUserIdFromToken(userToken)

	user := pc.userService.GetByID(userID)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success get profile", user))
}

func (pc *ProfileController) UpdateProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)

	userID := getUserIdFromToken(userToken)

	input := dto.UserRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", map[string]any{}))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", map[string]any{}))
	}

	if _, err := time.Parse("2006-01-02", input.BirthDate); err != nil && input.BirthDate != "" {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid, date format invalid", nil))
	}

	photo, err := c.FormFile("photo")

	if err == nil {
		url := helpers.UploadImage(photo, "pp")
		if url != "" {
			input.Photo = url
		}
	}

	user := pc.userService.Update(userID, input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success update profile", user))
}

func (pc *ProfileController) ChangePassword(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)

	userID := getUserIdFromToken(userToken)

	var passwords dto.ChangePasswordRequest

	if err := c.Bind(&passwords); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if err := passwords.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	err := pc.userService.ChangePassword(userID, passwords)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, fmt.Sprint(err), nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Successfully change password ", nil))
}

func getUserIdFromToken(userToken *jwt.Token) string {
	claims := userToken.Claims.(jwt.MapClaims)

	userIdInt := int(claims["id"].(float64))
	userID := strconv.Itoa(userIdInt)

	return userID
}
