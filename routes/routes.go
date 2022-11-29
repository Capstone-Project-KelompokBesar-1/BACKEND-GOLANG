package routes

import (
	"net/http"
	"ourgym/controllers"
	"ourgym/controllers/admin/manage/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	AuthController      controllers.AuthController
	AdminUserController users.AdminUserController
}

func (cl ControllerList) InitRoute() *echo.Echo {
	e := echo.New()

	e.POST("/login", cl.AuthController.Login)
	e.POST("/register", cl.AuthController.Register)

	admin := e.Group("/admin/users")
	admin.GET("", cl.AdminUserController.GetAll)
	admin.GET("/:id/:name", cl.AdminUserController.GetOneByFilter)
	admin.POST("", cl.AdminUserController.Create)
	admin.PUT("/:id", cl.AdminUserController.Update)
	admin.DELETE("/:id", cl.AdminUserController.Delete)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
