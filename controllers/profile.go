package controllers

import (
	"net/http"
	"ourgym/helpers"
	"ourgym/models"
	"ourgym/services"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success get profile", user.ConvertToDTO()))
}

func (pc *ProfileController) UpdateProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)

	userID := getUserIdFromToken(userToken)

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

	user := pc.userService.Update(userID, input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success update profile", user.ConvertToDTO()))
}

func (pc *ProfileController) ChangePassword(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)

	userID := getUserIdFromToken(userToken)

	var passwords models.ChangePassword

	if err := c.Bind(&passwords); err != nil {
		return c.JSON(http.StatusNotFound, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if err := passwords.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	user := pc.userService.GetByID(userID)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwords.OldPassword))

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "old password invalid", nil))
	}

	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwords.NewPassword), bcrypt.DefaultCost)

	isSuccess := pc.userService.ChangePassword(userID, string(newPassword))

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, Response(http.StatusInternalServerError, "failed to change password", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Successfully change password ", nil))
}

func getUserIdFromToken(userToken *jwt.Token) string {
	claims := userToken.Claims.(jwt.MapClaims)

	userIdInt := int(claims["id"].(float64))
	userID := strconv.Itoa(userIdInt)

	return userID
}
