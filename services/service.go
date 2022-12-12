package services

import (
	"ourgym/dto"
	"ourgym/models"

	"github.com/midtrans/midtrans-go/snap"
)

type AuthService interface {
	Login(userRequest models.User) (map[string]string, error)
	Register(userRequest models.User) error
	SendOTP(email string) error
	CreateNewPassword(otp, new_password string) error
}

type UserService interface {
	GetAll(name string) []dto.UserResponse
	GetByID(id string) dto.UserResponse
	Create(userRequest models.User) dto.UserResponse
	Update(id string, userRequest models.User) dto.UserResponse
	ChangePassword(id string, passwords dto.ChangePasswordRequest) error
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type ClassService interface {
	GetAll(classType string, name string) []dto.ClassResponse
	GetByID(id string) dto.ClassResponse
	Create(classRequest models.Class) dto.ClassResponse
	Update(id string, classRequest models.Class) dto.ClassResponse
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type TransactionService interface {
	GetAll() []dto.TransactionResponse
	GetHistory() []dto.TransactionResponse
	GetByUserID(userID string) []dto.TransactionResponse
	GetByID(id string) dto.TransactionResponse
	Create(transactionRequest dto.TransactionRequest) (snap.Response, error)
	UpdatedByMidtransAPI(midtransTransaction dto.MidtransTransactionRequest) error
	Update(id string, transactionRequest dto.TransactionRequest) dto.TransactionResponse
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type PaymentMethodService interface {
	GetAll() []dto.PaymentMethodResponse
	GetByID(id string) dto.PaymentMethodResponse
}

type TrainerService interface {
	GetAll(name string) []dto.TrainerResponse
	GetByID(id string) dto.TrainerResponse
}
