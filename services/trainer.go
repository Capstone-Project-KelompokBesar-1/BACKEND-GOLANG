package services

import (
	"ourgym/models"
	"ourgym/repositories"
)

func NewTrainerService(tr repositories.TrainerRepository) TrainerService {
	return &TrainerServiceImpl{
		trainerRepository: tr,
	}
}

type TrainerServiceImpl struct {
	trainerRepository repositories.TrainerRepository
}

func (ts *TrainerServiceImpl) GetAll(name string) []models.Trainer {
	return ts.trainerRepository.GetAll(name)
}

func (ts *TrainerServiceImpl) GetByID(id string) models.Trainer {
	return ts.trainerRepository.GetByID(id)
}
