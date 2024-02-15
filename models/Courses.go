package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model

	ID     int    `json:"id"`
	Name   string `json:"name"`
	Image  string `gorm:"type:varchar(255)"`
	UserID uint
}
