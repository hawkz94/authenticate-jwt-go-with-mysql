package models

import (
	"time"

	"gorm.io/gorm"
)

type Companies struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Photo     string `json:"photo"`
	IsDeleted bool   `json:"is_deleted"`

	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}
