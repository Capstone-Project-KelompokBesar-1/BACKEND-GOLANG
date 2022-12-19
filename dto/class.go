package dto

import "github.com/go-playground/validator/v10"

type ClassResponse struct {
	ID           uint             `json:"id"`
	TrainerID    uint             `json:"trainer_id"`
	Trainer      TrainerResponse  `json:"trainer"`
	CategoryID   uint             `json:"category_id"`
	Category     CategoryResponse `json:"category"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	TotalMeeting int              `json:"total_meeting"`
	Thumbnail    string           `json:"thumbnail"`
	Type         string           `json:"type"`
	Price        int              `json:"price"`
}

type ClassForTransactionResponse struct {
	ID           uint   `json:"id"`
	TrainerID    uint   `json:"trainer_id"`
	CategoryID   uint   `json:"category_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	TotalMeeting int    `json:"total_meeting"`
	Thumbnail    string `json:"thumbnail"`
	Type         string `json:"type"`
	Price        int    `json:"price"`
}

type ClassRequest struct {
	TrainerID    uint   `json:"trainer_id" form:"trainer_id" validate:"required"`
	CategoryID   uint   `json:"category_id" form:"category_id" validate:"required"`
	Name         string `json:"name" form:"name" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	TotalMeeting int    `json:"total_meeting" form:"total_meeting"`
	Thumbnail    string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Type         string `json:"type" form:"type" validate:"required"`
	Price        int    `json:"price" form:"price" validate:"required"`
}

func (cr *ClassRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(cr)

	return err
}
