package models

import (
	"ourgym/dto"
	"time"
)

type Trainer struct {
	ID          uint      `json:"id" form:"id" gorm:"primaryKey"`
	Name        string    `json:"name" form:"name"`
	Gender      string    `json:"gender" form:"gender"`
	Photo       string    `json:"photo" form:"photo"`
	Expertise   string    `json:"expertise" form:"expertise"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
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
