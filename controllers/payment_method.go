package controllers

import (
	"net/http"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type PaymentMethodController struct {
	paymentMethodService services.PaymentMethodService
}

func NewPaymentMethodController(paymentMethodService services.PaymentMethodService) *PaymentMethodController {
	return &PaymentMethodController{
		paymentMethodService,
	}
}

func (tc *PaymentMethodController) GetAll(c echo.Context) error {
	paymentMethods := tc.paymentMethodService.GetAll()

	if len(paymentMethods) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "payment methods not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get payment methods", paymentMethods))
}

func (tc *PaymentMethodController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	paymentMethod := tc.paymentMethodService.GetByID(id)

	if paymentMethod.ID == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "payment method not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "payment method found", paymentMethod))
}
