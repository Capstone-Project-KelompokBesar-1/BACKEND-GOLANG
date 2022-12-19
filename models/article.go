package models

import "time"

type Article struct {
	ID        string `gorm:"primaryKey"`
	Title     string
	Thumbnail string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
