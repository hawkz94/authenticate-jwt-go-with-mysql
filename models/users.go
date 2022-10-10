package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	ReferalCode   string `json:"referal_code"`
	TokenPassword string `json:"token_password"`
	IsActive      bool   `json:"is_active"`
	IsStaff       bool   `json:"is_staff"`
	IsDeleted     bool   `json:"is_deleted"`

	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
