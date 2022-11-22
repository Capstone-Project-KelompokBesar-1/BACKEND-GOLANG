package models

import (
	"time"

	"gorm.io/gorm"
)

type PhotoProfile struct {
	ID        uint           `json:"id" form:"id" gorm:"primaryKey"`
	UserId    uint           `json:"user_id" form:"user_id"`
	User      User           `json:"user" form:"user"`
	ImageUrl  string         `json:"image_url" form:"image_url"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
