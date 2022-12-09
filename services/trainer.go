package services

import (
	"ourgym/dto"
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

func (ts *TrainerServiceImpl) GetAll(name string) []dto.TrainerResponse {
	trainers := ts.trainerRepository.GetAll(name)

	var trainersResponse []dto.TrainerResponse

	for _, trainer := range trainers {
		trainersResponse = append(trainersResponse, trainer.ConvertToDTO())
	}

	return trainersResponse
}

func (ts *TrainerServiceImpl) GetByID(id string) dto.TrainerResponse {
	trainer := ts.trainerRepository.GetByID(id)
	return trainer.ConvertToDTO()
}
