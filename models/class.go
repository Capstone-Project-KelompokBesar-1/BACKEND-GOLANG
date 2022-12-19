package models

import (
	"ourgym/dto"
	"time"
)

type Class struct {
	ID           uint `gorm:"primaryKey"`
	TrainerID    uint
	Trainer      Trainer
	CategoryID   uint
	Category     Category
	Name         string
	Description  string
	TotalMeeting int
	Thumbnail    string
	Type         string
	Price        int
	Transactions []Transaction
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
		Name:         c.Name,
		Description:  c.Description,
		TotalMeeting: c.TotalMeeting,
		Thumbnail:    c.Thumbnail,
		Type:         c.Type,
		Price:        c.Price,
	}
}

func FromClassRequestToClassModel(request dto.ClassRequest) Class {
	return Class{
		TrainerID:    request.TrainerID,
		CategoryID:   request.CategoryID,
		Name:         request.Name,
		Description:  request.Description,
		TotalMeeting: request.TotalMeeting,
		Thumbnail:    request.Thumbnail,
		Type:         request.Type,
		Price:        request.Price,
	}
}
