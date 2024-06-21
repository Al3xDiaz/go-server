package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model

	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	Image  string `gorm:"type:varchar(255)" json:"image"`
	UserID uint   `json:"-"`
}
