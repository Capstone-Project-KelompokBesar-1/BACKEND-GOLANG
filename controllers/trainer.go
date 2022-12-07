package controllers

import (
	"net/http"
	"ourgym/dto"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type TrainerController struct {
	trainerService services.TrainerService
}

func NewTrainerController(trainerService services.TrainerService) *TrainerController {
	return &TrainerController{
		trainerService,
	}
}

func (tc *TrainerController) GetAll(c echo.Context) error {
	name := c.QueryParam("name")

	trainersData := tc.trainerService.GetAll(name)

	if len(trainersData) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Trainers Not Found", nil))
	}

	trainers := []dto.DTOTrainer{}

	for _, trainer := range trainersData {
		trainers = append(trainers, trainer.ConvertToDTO())
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success Get Trainers", trainers))
}

func (tc *TrainerController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	trainer := tc.trainerService.GetByID(id)

	if trainer.ID == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Trainer Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Trainer Found", trainer.ConvertToDTO()))
}
