package models

import (
	"ourgym/dto"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Phone        string
	Email        string
	Password     string
	Address      string
	Gender       string
	BirthDate    time.Time
	Photo        string
	IsAdmin      bool
	Transactions []Transaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u User) ConvertToDTO() dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		Email:     u.Email,
		Address:   u.Address,
		Gender:    u.Gender,
		BirthDate: u.BirthDate.Format("2006-01-02"),
		Photo:     u.Photo,
		IsAdmin:   u.IsAdmin,
	}
}

func (u *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}

func FromUserRequestToUserModel(request dto.UserRequest) User {
	birthDate, _ := time.Parse("2006-01-02", request.BirthDate)
	return User{
		Name:      request.Name,
		Phone:     request.Phone,
		Email:     request.Email,
		Password:  request.Password,
		Address:   request.Address,
		Gender:    request.Gender,
		BirthDate: birthDate,
		Photo:     request.Photo,
	}
}
