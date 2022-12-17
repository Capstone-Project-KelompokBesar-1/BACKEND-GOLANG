package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionResponse struct {
	ID              string                      `json:"id"`
	UserID          uint                        `json:"user_id"`
	User            UserResponse                `json:"user"`
	ClassID         uint                        `json:"class_id"`
	Class           ClassForTransactionResponse `json:"class"`
	PaymentMethodID uint                        `json:"payment_method_id"`
	PaymentMethod   PaymentMethodResponse       `json:"payment_method"`
	Amount          int                         `json:"amount"`
	Status          string                      `json:"status"`
	ExpiredAt       string                      `json:"expired_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
}

type TransactionRequest struct {
	UserID          uint   `json:"user_id" form:"user_id" validate:"required"`
	ClassID         uint   `json:"class_id" form:"class_id" validate:"required"`
	PaymentMethodID uint   `json:"payment_method_id" form:"payment_method_id" validate:"required"`
	Amount          int    `json:"amount" form:"amount" validate:"required"`
	Status          string `json:"status" form:"status"`
}

func (tr *TransactionRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(tr)

	return err
}
