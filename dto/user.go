package dto

import (
	"time"
)

type DTOUser struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Address   string    `json:"address" `
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	Photo     string    `json:"photo"`
	IsAdmin   bool      `json:"is_admin"`
}
