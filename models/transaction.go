package models

import (
	"ourgym/dto"
	"time"
)

type Transaction struct {
	ID              string `gorm:"primaryKey"`
	UserID          uint
	User            User
	ClassID         uint
	Class           Class
	PaymentMethodID uint
	PaymentMethod   PaymentMethod
	Amount          int
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
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
