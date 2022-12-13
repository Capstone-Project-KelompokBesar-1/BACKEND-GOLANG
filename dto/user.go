package dto

import (
	"github.com/go-playground/validator/v10"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Photo     string `json:"photo"`
	IsAdmin   bool   `json:"is_admin"`
}

type UserRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Phone     string `json:"phone" form:"phone" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Password  string `json:"password" form:"password"`
	Address   string `json:"address" form:"address"`
	Gender    string `json:"gender" form:"gender"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	Photo     string `json:"photo" form:"photo"`
}

func (ur *UserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(ur)

	return err
}
