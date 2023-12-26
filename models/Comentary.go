package models

import (
	"gorm.io/gorm"
)

type Commentary struct {
	gorm.Model

	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID  uint   `json:"userId"`
	Comment string `json:"comment"`
	SiteId  uint   `json:"-"`
	Site    Site   `json:"site" gorm:"foreignKey:SiteId:association_foreignkey:id;foreignkey:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
