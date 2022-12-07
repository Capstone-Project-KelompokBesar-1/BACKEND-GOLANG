package models

import "github.com/go-playground/validator/v10"

type ChangePassword struct {
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}

func (cp *ChangePassword) Validate() error {
	validate := validator.New()

	err := validate.Struct(cp)

	return err
}
