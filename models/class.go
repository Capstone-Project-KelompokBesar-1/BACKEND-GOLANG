package models

import "time"

type Class struct {
	ID          uint      `json:"id" form:"id" gorm:"primaryKey"`
	RoomID      uint      `json:"room_id" form:"room_id"`
	Room        Room      `json:"room" form:"room"`
	TrainerID   uint      `json:"trainer_id" form:"trainer_id"`
	Trainer     Trainer   `json:"trainer" form:"trainer"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	StartDate   time.Time `json:"start_date" form:"start_date"`
	Thumbnail   string    `json:"thumbnail" form:"thumbnail"`
	Type        string    `json:"type" form:"type"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}
