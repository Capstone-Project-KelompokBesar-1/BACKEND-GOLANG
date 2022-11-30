package user

import (
	"net/http"
	"ourgym/controllers"
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

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success get profile", user.ConvertToDTO()))
}

func (uc *ProfileController) UpdateProfile(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	userIdInt := int(claims["id"].(float64))
	userId := strconv.Itoa(userIdInt)

	input := models.User{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Failed to upload", map[string]any{}))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Request invalid", ""))
	}

	user := uc.userService.Update(userId, input)

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success update profile", user.ConvertToDTO()))
}

func (uc *ProfileController) UploadPhoto(c echo.Context) error {
	photo, err := c.FormFile("photo")

	if err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Failed to upload photo", ""))
	}

	photoUrl := helpers.UploadImage(photo)

	if photoUrl == "" {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Failed to upload photo", ""))
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success upload photo", map[string]string{
		"url_photo": photoUrl,
	}))
}
