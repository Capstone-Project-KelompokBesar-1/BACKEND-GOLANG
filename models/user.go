package models

import (
	"ourgym/dto"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint      `json:"id" form:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name" validate:"required"`
	Phone     string    `json:"phone" form:"phone" validate:"required"`
	Email     string    `json:"email" form:"email" validate:"required,email"`
	Password  string    `json:"password" form:"password"`
	Address   string    `json:"address" form:"address"`
	Gender    string    `json:"gender" form:"gender"`
	BirthDate time.Time `json:"birth_date" form:"birth_date"`
	Photo     string    `json:"photo" form:"photo"`
	IsAdmin   bool      `json:"is_admin" form:"is_admin"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func (u User) ConvertToDTO() dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		Email:     u.Email,
		Address:   u.Address,
		Gender:    u.Gender,
		BirthDate: u.BirthDate,
		Photo:     u.Photo,
		IsAdmin:   u.IsAdmin,
	}
}

func (u *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}
