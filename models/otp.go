package models

import "time"

type Otp struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Code      int
	ExpiredAt time.Time
}
