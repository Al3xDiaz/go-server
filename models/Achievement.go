package models

import "gorm.io/gorm"

type AchievementsHistory struct {
	gorm.Model

	ID      uint   `gorm:"primaryKey" json:"id"`
	Year    int    `json:"year" gorm:"check:year>0"`
	Comment string `json:"comment" gorm:"check:comment<>''"`
	Title   string `json:"title" gorm:"check:title<>''"`
	UserID  uint   `json:"-"`
}
