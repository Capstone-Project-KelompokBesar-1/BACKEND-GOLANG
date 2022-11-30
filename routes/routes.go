package routes

import (
	"net/http"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/controllers/admin/manage/users"
	"ourgym/controllers/user"
	"ourgym/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController    controllers.AuthController
	UserController    users.UserController
	ProfileController user.ProfileController
}

func (cl ControllerList) InitRoute() *echo.Echo {
	e := echo.New()
	cfg := config.Cfg

	e.POST("/login", cl.AuthController.Login)
	e.POST("/register", cl.AuthController.Register)

	user := e.Group("")

	user.Use(middleware.JWT([]byte(cfg.JWT_SECRET_KEY)))

	user.GET("user/profile", cl.ProfileController.GetProfile)
	user.PUT("user/profile", cl.ProfileController.UpdateProfile)
	user.POST("/upload-photo", cl.ProfileController.UploadPhoto)
	user.POST("/refresh-token", cl.AuthController.RefreshToken)

	configAdmin := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForAdmin,
	}

	admin := e.Group("")

	admin.Use(middleware.JWTWithConfig(configAdmin))

	admin.GET("/users", cl.UserController.GetAll)
	admin.GET("/users/:id", cl.UserController.GetOneByFilter)
	admin.POST("/users", cl.UserController.Create)
	admin.PUT("/users/:id", cl.UserController.Update)
	admin.DELETE("/users/:id", cl.UserController.Delete)
	admin.DELETE("/users", cl.UserController.DeleteMany)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
