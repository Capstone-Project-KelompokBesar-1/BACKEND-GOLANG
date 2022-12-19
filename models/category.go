package models

import (
	"ourgym/dto"
)

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (c Category) ConvertToDTO() dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}
