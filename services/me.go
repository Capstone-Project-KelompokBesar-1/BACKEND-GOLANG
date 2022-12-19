package services

import (
	"ourgym/dto"
	"ourgym/repositories"
)

func NewMeService(me repositories.MeRepository) MeService {
	return &MeServiceImpl{
		meRepository: me,
	}
}

type MeServiceImpl struct {
	meRepository repositories.MeRepository
}

func (me *MeServiceImpl) OnlineClass(userId string) []dto.ClassForTransactionResponse {
	transactions := me.meRepository.GetTransactionByID(userId)

	var onlineClassResponse []dto.ClassForTransactionResponse

	for _, transaction := range transactions {
		if transaction.Status == "settlement" && transaction.Class.Type == "online" {
			onlineClassResponse = append(onlineClassResponse, transaction.ConvertToDTO().Class)
		}
	}

	return onlineClassResponse
}

func (me *MeServiceImpl) OfflineClass(userId string) []dto.ClassForTransactionResponse {
	transactions := me.meRepository.GetTransactionByID(userId)

	var offlineClassResponse []dto.ClassForTransactionResponse

	for _, transaction := range transactions {
		if transaction.Status == "settlement" && transaction.Class.Type == "offline" {
			offlineClassResponse = append(offlineClassResponse, transaction.ConvertToDTO().Class)
		}
	}

	return offlineClassResponse
}
