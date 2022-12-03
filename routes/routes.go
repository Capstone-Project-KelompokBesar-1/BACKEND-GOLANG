package routes

import (
	"net/http"
	"ourgym/config"
	"ourgym/controllers"
	"ourgym/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	AuthController    controllers.AuthController
	UserController    controllers.UserController
	ProfileController controllers.ProfileController
	ClassController   controllers.ClassController
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

	user.GET("/classes", cl.ClassController.GetAll)
	user.GET("/classes/online", cl.ClassController.GetAllOnlineClass)
	user.GET("/classes/offline", cl.ClassController.GetAllOfflineClass)
	admin.GET("/classes/:id", cl.ClassController.GetByID)
	admin.POST("/classes", cl.ClassController.Create)
	admin.PUT("/classes/:id", cl.ClassController.Update)
	admin.DELETE("/classes/:id", cl.ClassController.Delete)
	admin.DELETE("/classes", cl.ClassController.DeleteMany)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
