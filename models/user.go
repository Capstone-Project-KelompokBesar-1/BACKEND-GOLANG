package models

import (
	"ourgym/dto"
	"time"
)

type User struct {
	ID uint `json:"id" form:"id" gorm:"primaryKey"`
	// ProfileID uint      `json:"profile_id" form:"profile_id"`
	Profile   Profile   `json:"profile" form:"profile"`
	Name      string    `json:"name" form:"name"`
	Phone     string    `json:"phone" form:"phone"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	IsAdmin   bool      `json:"is_admin" form:"is_admin"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func (u User) ConvertToDTO() dto.DTOUser {
	return dto.DTOUser{
		ID:      u.ID,
		Name:    u.Name,
		Phone:   u.Phone,
		Email:   u.Email,
		IsAdmin: u.IsAdmin,
		Address: u.Profile.Address,
	}
}
