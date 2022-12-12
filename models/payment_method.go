package models

import "ourgym/dto"

type PaymentMethod struct {
	ID   uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
	Code string `json:"code" form:"code"`
}

func (pm PaymentMethod) ConvertToDTO() dto.PaymentMethodResponse {
	return dto.PaymentMethodResponse{
		ID:   pm.ID,
		Name: pm.Name,
	}
}
