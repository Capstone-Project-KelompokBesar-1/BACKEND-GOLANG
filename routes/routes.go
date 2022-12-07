package routes

import (
	"net/http"
	"ourgym/controllers"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	AuthController controllers.AuthController
}

func (cl ControllerList) InitRoute() *echo.Echo {
	e := echo.New()

	e.POST("/login", cl.AuthController.Login)
	e.POST("/register", cl.AuthController.Register)
	e.POST("/send-otp", cl.AuthController.SendOTP)
	e.POST("/forgot-password", cl.AuthController.CreateNewPassword)

	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
