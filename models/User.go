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

type User struct {
	gorm.Model

	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserName string `json:"userName" gorm:"type=varchar(100);unique;not null"`
	Email    string `json:"email" gorm:"type:varchar(50);not null"`
	Password string `json:"password,omitempty" gorm:"size:255"`
	Verified bool   `json:"verified" gorm:"default=false"`

	Permisions []*Permision `gorm:"many2many:UserPermision;association_foreignkey:id;foreignkey:id"`
	Profile    Profile      `json:"profile"`
}
type Profile struct {
	gorm.Model

	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"userId" gorm:"not null"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Telephone string `json:"telephone" gorm:"null"`
	Photo     string `json:"photo" gorm:"null"`

	// Social Media
	Linkedin  string `json:"linkedin" gorm:"null"`
	Github    string `json:"github" gorm:"null"`
	Twitter   string `json:"twitter" gorm:"null"`
	Facebook  string `json:"facebook" gorm:"null"`
	Instagram string `json:"instagram" gorm:"null"`
	Youtube   string `json:"youtube" gorm:"null"`
	Website   string `json:"website" gorm:"null"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash := MakePassword(u.Password)
		tx.Statement.SetColumn("Password", hash)
	}

	fmt.Println(u.Password)
	return
}

type Permision struct {
	gorm.Model

	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string `json:"userName" gorm:"type=varchar(150)"`
	Path        string `json:"path" gorm:"type=varchar(50)"`    // /*, /api/users/*, /api/users/{id}
	Methods     string `json:"methods" gorm:"type=varchar(50)"` // (*), (get|post), (get|post|put|delete)
	// Users       []*User `gorm:"many2many:user_permision;"`
}

func MakeMapString(list []*Permision) []string {
	var resp = make([]string, 0)
	for _, v := range list {
		resp = append(resp, v.Path)
	}
	return resp
}
