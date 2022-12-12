package dto

type PaymentMethodResponse struct {
	ID   uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name" validate:"required"`
}
