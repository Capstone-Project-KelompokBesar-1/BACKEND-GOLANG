package services

import (
	"ourgym/dto"

	"github.com/midtrans/midtrans-go/snap"
)

type AuthService interface {
	Login(loginRequest dto.LoginRequest) (map[string]string, error)
	Register(userRequest dto.UserRequest) error
	ForgotPassword(email string) error
	ValidateOTP(otpCode int) (map[string]string, error)
	ResetPassword(otp, new_password string) error
}

type UserService interface {
	GetAll(name string) []dto.UserResponse
	GetByID(id string) dto.UserResponse
	Create(userRequest dto.UserRequest) (dto.UserResponse, error)
	Update(id string, userRequest dto.UserRequest) dto.UserResponse
	ChangePassword(id string, passwords dto.ChangePasswordRequest) error
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type ClassService interface {
	GetAll(classType string, name string) []dto.ClassResponse
	GetByID(id string) dto.ClassResponse
	Create(classRequest dto.ClassRequest) dto.ClassResponse
	Update(id string, classRequest dto.ClassRequest) dto.ClassResponse
	Delete(id string) bool
	DeleteMany(ids string) bool
}

type TransactionService interface {
	GetAll() []dto.TransactionResponse
	GetHistory() []dto.TransactionResponse
	GetByUserID(userID, status string) []dto.TransactionResponse
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

type MeService interface {
	OnlineClass(userId string) []dto.ClassForTransactionResponse
	OfflineClass(userId string) []dto.ClassForTransactionResponse
}

type CategoryService interface {
	GetAll(name string) []dto.CategoryResponse
	GetByID(id string) dto.CategoryResponse
}

type DashboardService interface {
	GetData() dto.DashboardResponse
}
