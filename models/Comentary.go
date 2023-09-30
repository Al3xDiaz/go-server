package models

import (
	"gorm.io/gorm"
)	

type Commentary struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Email string `json:"email" gorm:"type:varchar(50);not null"`
	Comment string	`json:"comment"`
}

type Commentaries []Commentary