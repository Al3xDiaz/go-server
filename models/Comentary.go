package models

import (
	"gorm.io/gorm"
)

type Commentary struct {
	gorm.Model

	ID      uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	UserID  uint   `json:"userId"`
	Comment string `json:"comment"`
}
