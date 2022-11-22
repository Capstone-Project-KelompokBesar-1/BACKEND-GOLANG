package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Success get data",
			"data":    "Hello world!",
		})
	})

	return e
}
