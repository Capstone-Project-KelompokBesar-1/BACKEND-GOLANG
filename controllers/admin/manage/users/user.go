package users

import (
	"net/http"
	"ourgym/controllers"
	"ourgym/dto"
	"ourgym/models"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type AdminUserController struct {
	adminUserService services.UserServices
}

func NewAdminUserController(adminUserService services.UserServices) *AdminUserController {
	return &AdminUserController{
		adminUserService,
	}
}

func (uc *AdminUserController) GetAll(c echo.Context) error {
	usersData := uc.adminUserService.GetAll()

	users := []dto.DTOUser{}

	for _, user := range usersData {
		users = append(users, user.ConvertToDTO())
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success GetAll User", users))
}

func (uc *AdminUserController) GetOneByFilter(c echo.Context) error {
	var id string = c.Param("id")
	var name string = c.Param("name")

	user := uc.adminUserService.GetOneByFilter(id, name)

	if user.ID == 0 || user.Name == "" {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusNotFound, "User Not Found", ""))
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "User Found", user.ConvertToDTO()))
}

func (uc *AdminUserController) Create(c echo.Context) error {
	input := models.User{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, controllers.Response(http.StatusBadRequest, "Failed", ""))
	}

	user := uc.adminUserService.Create(input)

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success Created User", user))
}

func (uc *AdminUserController) Update(c echo.Context) error {
	input := models.User{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusBadRequest, "Failed", ""))
	}

	var userId string = c.Param("id")

	user := uc.adminUserService.Update(userId, input)

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "Success Update User", user.ConvertToDTO()))
}

func (uc *AdminUserController) Delete(c echo.Context) error {
	var userId string = c.Param("id")

	isSuccess := uc.adminUserService.Delete(userId)

	if isSuccess {
		return c.JSON(http.StatusNotFound, controllers.Response(http.StatusNotFound, "User Not Found", ""))
	}

	return c.JSON(http.StatusOK, controllers.Response(http.StatusOK, "User Success Deleted", ""))
}
