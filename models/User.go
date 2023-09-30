package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"gorm.io/gorm"
)

func MakePassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return hex.EncodeToString(sum[:])
}

type Users struct {
	gorm.Model

	ID         uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName   string       `json:"userName" gorm:"type=varchar(100);unique;not null"`
	Email      string       `json:"email" gorm:"type:varchar(50);not null"`
	Password   string       `json:"password,omitempty" gorm:"size:255"`
	FirstName  string       `json:"firstName" gorm:"not null"`
	LastName   string       `json:"lastName" gorm:"not null"`
	Verified   bool         `json:"verified" gorm:"default=false"`
	Permisions []Permisions `json:"permisions" gorm:"foreignKey:UserID"`
}

func (u *Users) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash := MakePassword(u.Password)
		tx.Statement.SetColumn("Password", hash)
	}

	fmt.Println(u.Password)
	return
}

type Permisions struct {
	gorm.Model

	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string `json:"userName" gorm:"type=varchar(150)"`
	Path        string `json:"path" gorm:"type=varchar(50)"`    // /*, /api/users/*, /api/users/{id}
	Methods     string `json:"methods" gorm:"type=varchar(50)"` // (*), (get|post), (get|post|put|delete)
	UserID      uint   `json:"userId" gorm:"foreignKey:UserID"`
}
