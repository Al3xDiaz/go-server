package models

import (
	"gorm.io/gorm"
)

type Gallery struct {
	gorm.Model

	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement,not null"`
	Image  string `json:"image" gorm:"type=nvarchar(100);not null"`
	UserID uint   `json:"-" gorm:"null"`
}
