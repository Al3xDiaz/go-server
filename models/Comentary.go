package models

import (
	"gorm.io/gorm"
)

type Commentary struct {
	gorm.Model

	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement;"`
	UserID  uint   `json:"userId"`
	Comment string `json:"comment"`
	SiteId  uint   `json:"-" gorm:"null"`
	Site    Site   `json:"site" gorm:"foreignKey:SiteId:association_foreignkey:id;foreignkey:site_id;constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}
