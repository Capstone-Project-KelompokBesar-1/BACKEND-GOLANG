package controllers

import (
	"net/http"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type MeController struct {
	meService services.MeService
}

func NewMeController(meService services.MeService) *MeController {
	return &MeController{
		meService,
	}
}

func (me *MeController) OnlineClass(c echo.Context) error {
	var userID string = c.Param("id")

	onlineClass := me.meService.OnlineClass(userID)

	if len(onlineClass) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "classes not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get online classes by user", onlineClass))
}

func (me *MeController) OfflineClass(c echo.Context) error {
	var userID string = c.Param("id")

	offlineClass := me.meService.OfflineClass(userID)

	if len(offlineClass) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "classes not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get offline classes by user", offlineClass))
}
