package models

type Category struct {
	ID   uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
}
