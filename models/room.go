package models

type Room struct {
	ID           uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name         string `json:"name" form:"name"`
	Url          string `json:"url" form:"url"`
	Address      string `json:"address" form:"address"`
	Requirements string `json:"requirements" form:"requirements"`
}
