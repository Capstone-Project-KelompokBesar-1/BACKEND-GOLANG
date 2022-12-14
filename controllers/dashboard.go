package controllers

import (
	"net/http"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type DashboardController struct {
	dashboardService services.DashboardService
}

func NewDashboardController(dashboardService services.DashboardService) *DashboardController {
	return &DashboardController{
		dashboardService,
	}
}

func (tc *DashboardController) GetData(c echo.Context) error {
	dashboardData := tc.dashboardService.GetData()

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get dashboard data", dashboardData))
}
