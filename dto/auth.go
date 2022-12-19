package dto

import "github.com/go-playground/validator/v10"

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}

func (cp *ChangePasswordRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(cp)

	return err
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (lr *LoginRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(lr)

	return err
}
