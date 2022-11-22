package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `json:"id" form:"id" gorm:"primaryKey"`
	Name      string         `json:"name" form:"name"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
