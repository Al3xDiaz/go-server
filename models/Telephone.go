package models

import "gorm.io/gorm"

type Telephone struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey" json:"id"`
	ProfileID   uint   `json:"profileID"`
	PhoneNumber string `gorm:"size:20" json:"phoneNumber"`
	Whatsapp    bool   `gorm:"default:false" json:"whatsapp"`
	CountryCode string `gorm:"size:3" json:"countryCode"`
}
