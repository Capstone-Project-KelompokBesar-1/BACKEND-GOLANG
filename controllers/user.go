package controllers

import (
	"fmt"
	"net/http"
	"ourgym/dto"
	"ourgym/helpers"
	"ourgym/services"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) GetAll(c echo.Context) error {
	name := c.QueryParam("name")

	users := uc.userService.GetAll(name)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success Get Users", users))
}

func (uc *UserController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	user := uc.userService.GetByID(id)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "User Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "User Found", user))
}

func (uc *UserController) Create(c echo.Context) error {
	input := dto.UserRequest{}

	if err := c.Bind(&input); err != nil || input.Password == "" {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
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

	user, err := uc.userService.Create(input)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, fmt.Sprint(err), nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success Created User", user))
}

func (uc *UserController) Update(c echo.Context) error {
	input := dto.UserRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
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

	var userId string = c.Param("id")

	user := uc.userService.Update(userId, input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success Update User", user))
}

func (uc *UserController) Delete(c echo.Context) error {
	var userId string = c.Param("id")

	isSuccess := uc.userService.Delete(userId)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "User Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "User Success Deleted", nil))
}

func (uc *UserController) DeleteMany(c echo.Context) error {
	ids := c.QueryParam("ids")

	isSuccess := uc.userService.DeleteMany(ids)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Users Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Users Success Deleted", nil))
}
