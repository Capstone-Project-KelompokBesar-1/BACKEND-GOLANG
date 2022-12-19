package services

import (
	"ourgym/dto"
	"ourgym/repositories"
)

func NewPaymentMethodService(tr repositories.PaymentMethodRepository) PaymentMethodService {
	return &PaymentMethodServiceImpl{
		paymentMethodRepository: tr,
	}
}

type PaymentMethodServiceImpl struct {
	paymentMethodRepository repositories.PaymentMethodRepository
}

func (ts *PaymentMethodServiceImpl) GetAll() []dto.PaymentMethodResponse {
	paymentMethods := ts.paymentMethodRepository.GetAll()

	var paymentMethodsResponse []dto.PaymentMethodResponse

	for _, paymentMethod := range paymentMethods {
		paymentMethodsResponse = append(paymentMethodsResponse, paymentMethod.ConvertToDTO())
	}

	return paymentMethodsResponse
}

func (ts *PaymentMethodServiceImpl) GetByID(id string) dto.PaymentMethodResponse {
	paymentMethod := ts.paymentMethodRepository.GetByID(id)
	return paymentMethod.ConvertToDTO()
}
