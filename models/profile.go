package models

import "time"

type Profile struct {
	ID        uint      `json:"id" form:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" form:"user_id"`
	Address   string    `json:"address" form:"address"`
	Gender    string    `json:"gender" form:"gender"`
	BirthDate time.Time `json:"birth_date" form:"birth_date"`
	Photo     string    `json:"photo" form:"photo"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
