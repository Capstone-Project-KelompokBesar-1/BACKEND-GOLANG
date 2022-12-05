package controllers

import (
	"net/http"
	"ourgym/helpers"
	"ourgym/models"
	"ourgym/services"
	"strconv"

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

func (uc *ProfileController) GetProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	userIdInt := int(claims["id"].(float64))
	userID := strconv.Itoa(userIdInt)

	user := uc.userService.GetByID(userID)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success get profile", user.ConvertToDTO()))
}

func (uc *ProfileController) UpdateProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	userIdInt := int(claims["id"].(float64))
	userId := strconv.Itoa(userIdInt)

	input := models.User{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", map[string]any{}))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", map[string]any{}))
	}

	photo, err := c.FormFile("photo")

	if err == nil {
		url := helpers.UploadImage(photo, "pp")
		if url != "" {
			input.Photo = url
		}
	}

	user := uc.userService.Update(userId, input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success update profile", user.ConvertToDTO()))
}
