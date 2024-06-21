package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement,not null"`
	Title       string `json:"title" gorm:"type=varchar(100);not null"`
	Description string `json:"description" gorm:"type=nvarchar(300);not null"`
	Image       string `json:"image" gorm:"type=nvarchar(100);null"`
	URL         string `json:"url" gorm:"null"`
	UserID      uint   `json:"-" gorm:"null"`

	StartDate time.Time `json:"startDate" gorm:"not null"`
	EndDate   time.Time `json:"endDate" gorm:"null"`
}
