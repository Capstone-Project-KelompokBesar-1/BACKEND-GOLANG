package models

type Otp struct {
	ID    uint   `json:"id" form:"id" gorm:"primaryKey"`
	User  uint   `json:"user" form:"user"`
	Code  string `json:"code" form:"code"`
	Token string `json:"token" form:"token"`
}
