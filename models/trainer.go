package models

import (
	"ourgym/dto"
	"time"
)

type Trainer struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Gender      string
	Photo       string
	Expertise   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Trainer) ConvertToDTO() dto.TrainerResponse {
	return dto.TrainerResponse{
		ID:          t.ID,
		Name:        t.Name,
		Gender:      t.Gender,
		Photo:       t.Photo,
		Expertise:   t.Expertise,
		Description: t.Description,
	}
}
