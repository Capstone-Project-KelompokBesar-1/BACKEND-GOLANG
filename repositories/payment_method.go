package repositories

import (
	"ourgym/models"

	"gorm.io/gorm"
)

func NewPaymentMethodRepository(gormDB *gorm.DB) PaymentMethodRepository {
	return &PaymentMethodRepositoryImpl{
		db: gormDB,
	}
}

type PaymentMethodRepositoryImpl struct {
	db *gorm.DB
}

func (pr *PaymentMethodRepositoryImpl) GetAll() []models.PaymentMethod {
	var paymentMethods []models.PaymentMethod

	pr.db.Find(&paymentMethods)

	return paymentMethods
}

func (pr *PaymentMethodRepositoryImpl) GetByID(id string) models.PaymentMethod {
	var paymentMethod models.PaymentMethod

	pr.db.First(&paymentMethod, id)

	return paymentMethod
}
