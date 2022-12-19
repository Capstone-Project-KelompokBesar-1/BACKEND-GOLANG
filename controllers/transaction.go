package controllers

import (
	"log"
	"net/http"
	"ourgym/dto"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService,
	}
}

func (uc *TransactionController) GetAll(c echo.Context) error {
	transactions := uc.transactionService.GetAll()

	if len(transactions) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transactions not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get transactions", transactions))
}

func (uc *TransactionController) GetHistory(c echo.Context) error {
	transactions := uc.transactionService.GetHistory()

	if len(transactions) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transactions history not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get transactions history", transactions))
}

func (uc *TransactionController) GetByUserID(c echo.Context) error {
	var userID string = c.Param("id")
	var status string = c.QueryParam("status")

	transactions := uc.transactionService.GetByUserID(userID, status)

	if len(transactions) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transactions not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success get transactions by user id", transactions))
}

func (uc *TransactionController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	transaction := uc.transactionService.GetByID(id)

	if transaction.ID == "" {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transaction not found", nil))
	}

	if transaction.ID == "" {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transaction not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "transaction found", transaction))
}

func (uc *TransactionController) Create(c echo.Context) error {
	input := dto.TransactionRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "request invalid", nil))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "request invalid", nil))
	}

	snapRes, err := uc.transactionService.Create(input)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response(http.StatusInternalServerError, "can not create a transaction", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success Created Transaction", snapRes))
}

func (uc *TransactionController) UpdatedByMidtransAPI(c echo.Context) error {
	input := dto.MidtransTransactionRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "request invalid", nil))
	}

	err := uc.transactionService.UpdatedByMidtransAPI(input)

	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success update transaction by midtrans", nil))
}

func (uc *TransactionController) Update(c echo.Context) error {
	input := dto.TransactionRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "request invalid", nil))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "request invalid", nil))
	}

	var transactionId string = c.Param("id")

	transaction := uc.transactionService.Update(transactionId, input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "success update transaction", transaction))
}

func (uc *TransactionController) Delete(c echo.Context) error {
	var transactionId string = c.Param("id")

	isSuccess := uc.transactionService.Delete(transactionId)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transaction not nound", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "transaction success deleted", nil))
}

func (uc *TransactionController) DeleteMany(c echo.Context) error {
	ids := c.QueryParam("ids")

	isSuccess := uc.transactionService.DeleteMany(ids)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "transactions not found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "transactions success deleted", nil))
}
