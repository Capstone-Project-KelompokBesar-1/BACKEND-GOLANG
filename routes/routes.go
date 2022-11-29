package routes

import (
	"net/http"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/controllers/admin/manage/users"
	"ourgym/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController      controllers.AuthController
	AdminUserController users.AdminUserController
}

func (cl ControllerList) InitRoute() *echo.Echo {
	e := echo.New()
	cfg := config.Cfg

	e.POST("/login", cl.AuthController.Login)
	e.POST("/register", cl.AuthController.Register)

	user := e.Group("")

	user.Use(middleware.JWT([]byte(cfg.JWT_SECRET_KEY)))

	user.POST("/refresh-token", cl.AuthController.RefreshToken)

	configAdmin := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForAdmin,
	}

	admin := e.Group("")

	admin.Use(middleware.JWTWithConfig(configAdmin))

	admin.GET("/users", cl.AdminUserController.GetAll)
	admin.GET("/users/:id/:name", cl.AdminUserController.GetOneByFilter)
	admin.POST("/users", cl.AdminUserController.Create)
	admin.PUT("/users", cl.AdminUserController.Update)
	admin.DELETE("/users/:id", cl.AdminUserController.Delete)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
