package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" form:"id" gorm:"primaryKey"`
	RoleId    uint      `json:"role_id" form:"role_id"`
	Role      Role      `json:"role" form:"role"`
	Name      string    `json:"name" form:"name"`
	Phone     string    `json:"phone" form:"phone"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}
