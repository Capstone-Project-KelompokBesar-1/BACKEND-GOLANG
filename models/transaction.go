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
	var expiredAt string = t.CreatedAt.Local().Add(time.Hour * 24).String()
	if t.Status == "settlement" || t.Status == "capture" {
		t.Status = "berhasil"
		expiredAt = ""
	} else if t.Status == "pending" {
		t.Status = "tertunda"
	} else {
		t.Status = "gagal"
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
			BirthDate: t.User.BirthDate.Format("2006-01-02"),
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
		PaymentMethodID: t.PaymentMethodID,
		PaymentMethod: dto.PaymentMethodResponse{
			ID:   t.PaymentMethod.ID,
			Name: t.PaymentMethod.Name,
		},
		Amount:    t.Amount,
		Status:    t.Status,
		ExpiredAt: expiredAt,
		UpdatedAt: t.UpdatedAt.Local(),
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
