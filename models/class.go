package models

import (
	"time"

	"ourgym/dto"

	"github.com/go-playground/validator/v10"
)

type Class struct {
	ID          uint      `json:"id" form:"id" gorm:"primaryKey"`
	TrainerID   uint      `json:"trainer_id" form:"trainer_id" validate:"required"`
	Trainer     Trainer   `json:"trainer" form:"trainer"`
	CategoryID  uint      `json:"category_id" form:"category_id" validate:"required"`
	Category    Category  `json:"category" form:"category"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	Thumbnail   string    `json:"thumbnail" form:"thumbnail"`
	Type        string    `json:"type" form:"type" validate:"required"`
	Price       int       `json:"price" form:"price" validate:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

func (c Class) ConvertToDTO() dto.ClassResponse {
	return dto.ClassResponse{
		ID:        c.ID,
		TrainerID: c.TrainerID,
		Trainer: dto.TrainerResponse{
			ID:          c.Trainer.ID,
			Name:        c.Trainer.Name,
			Gender:      c.Trainer.Gender,
			Photo:       c.Trainer.Photo,
			Expertise:   c.Trainer.Expertise,
			Description: c.Trainer.Description,
		},
		CategoryID: c.CategoryID,
		Category: dto.CategoryResponse{
			ID:   c.Category.ID,
			Name: c.Category.Name,
		},
		Name:        c.Name,
		Description: c.Description,
		Thumbnail:   c.Thumbnail,
		Type:        c.Type,
		Price:       c.Price,
	}
}

func (c *Class) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	return err
}
