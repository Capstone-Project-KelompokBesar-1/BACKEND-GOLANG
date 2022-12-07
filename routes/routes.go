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

	configAdmin := middleware.JWTConfig{
		KeyFunc: middlewares.GetJWTSecretKeyForAdmin,
	}

	adminJwtMiddleware := middleware.JWTWithConfig(configAdmin)

	userJwtMiddleware := middleware.JWT([]byte(cfg.JWT_SECRET_KEY))

	user := e.Group("/user")

	user.GET("/profile", cl.ProfileController.GetProfile, userJwtMiddleware)
	user.PUT("/profile", cl.ProfileController.UpdateProfile, userJwtMiddleware)
	user.PUT("/change-password", cl.ProfileController.ChangePassword, userJwtMiddleware)
	user.POST("user/refresh-token", cl.AuthController.RefreshToken, userJwtMiddleware)

	users := e.Group("/users")

	users.GET("", cl.UserController.GetAll, adminJwtMiddleware)
	users.GET("/:id", cl.UserController.GetOneByFilter, adminJwtMiddleware)
	users.POST("", cl.UserController.Create, adminJwtMiddleware)
	users.PUT("/:id", cl.UserController.Update, adminJwtMiddleware)
	users.DELETE("/:id", cl.UserController.Delete, adminJwtMiddleware)
	users.DELETE("", cl.UserController.DeleteMany, adminJwtMiddleware)

	classes := e.Group("/classes")

	classes.GET("", cl.ClassController.GetAll, userJwtMiddleware)
	classes.GET("/online", cl.ClassController.GetAllOnlineClass, userJwtMiddleware)
	classes.GET("/offline", cl.ClassController.GetAllOfflineClass, userJwtMiddleware)
	classes.GET("/:id", cl.ClassController.GetByID, userJwtMiddleware)
	classes.POST("", cl.ClassController.Create, adminJwtMiddleware)
	classes.PUT("/:id", cl.ClassController.Update, adminJwtMiddleware)
	classes.DELETE("/:id", cl.ClassController.Delete, adminJwtMiddleware)
	classes.DELETE("", cl.ClassController.DeleteMany, adminJwtMiddleware)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
