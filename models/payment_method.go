package models

import "ourgym/dto"

type PaymentMethod struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Code string
}

func (pm PaymentMethod) ConvertToDTO() dto.PaymentMethodResponse {
	return dto.PaymentMethodResponse{
		ID:   pm.ID,
		Name: pm.Name,
	}
}
