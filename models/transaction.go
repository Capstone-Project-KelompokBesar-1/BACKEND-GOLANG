package models

import (
	"ourgym/dto"
	"time"
)

type Transaction struct {
	ID              string        `json:"id" form:"id" gorm:"primaryKey"`
	UserID          uint          `json:"user_id" form:"user_id" validate:"required"`
	User            User          `json:"user" form:"user"`
	ClassID         uint          `json:"class_id" form:"class_id" validate:"required"`
	Class           Class         `json:"class" form:"class"`
	PaymentMethodID uint          `json:"payment_method_id" form:"payment_method_id" validate:"required"`
	PaymentMethod   PaymentMethod `json:"payment_method" form:"payment_method"`
	Amount          int           `json:"amount" form:"amount" validate:"required"`
	Status          string        `json:"status" form:"status"`
	CreatedAt       time.Time     `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" form:"updated_at"`
}

func (t Transaction) ConvertToDTO() dto.TransactionResponse {
	var expiredAt, successAt string
	if t.Status == "settlement" || t.Status == "capture" {
		t.Status = "berhasil"
		successAt = t.UpdatedAt.String()
	} else if t.Status == "pending" {
		t.Status = "tertunda"
		expiredAt = t.CreatedAt.Add(time.Hour * 24).String()
	} else {
		t.Status = "gagal"
		expiredAt = t.CreatedAt.Add(time.Hour * 24).String()
	}

	return dto.TransactionResponse{
		ID:     t.ID,
		UserID: t.UserID,
		User: dto.UserResponse{
			ID:        t.User.ID,
			Name:      t.User.Name,
			Phone:     t.User.Phone,
			Email:     t.User.Email,
			Address:   t.User.Address,
			Gender:    t.User.Gender,
			BirthDate: t.User.BirthDate,
			Photo:     t.User.Photo,
			IsAdmin:   t.User.IsAdmin,
		},
		ClassID: t.ClassID,
		Class: dto.ClassForTransactionResponse{
			ID:           t.Class.ID,
			TrainerID:    t.Class.TrainerID,
			CategoryID:   t.Class.CategoryID,
			Name:         t.Class.Name,
			Description:  t.Class.Description,
			TotalMeeting: t.Class.TotalMeeting,
			Thumbnail:    t.Class.Thumbnail,
			Type:         t.Class.Type,
			Price:        t.Class.Price,
		},
		PaymentMethod: t.PaymentMethod.Name,
		Amount:        t.Amount,
		Status:        t.Status,
		ExpiredAt:     expiredAt,
		SuccessAt:     successAt,
	}
}

func FromTransactionRequest(request dto.TransactionRequest) Transaction {
	return Transaction{
		UserID:          request.UserID,
		ClassID:         request.ClassID,
		PaymentMethodID: request.PaymentMethodID,
		Amount:          request.Amount,
		Status:          request.Status,
	}
}
