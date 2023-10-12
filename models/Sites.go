package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model

	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Url         string `json:"url" gorm:"type=varchar(100);unique;not null"`
	Description string `json:"description" gorm:"type=varchar(100)"`
}
