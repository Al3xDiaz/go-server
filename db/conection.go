package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
var DSN = "golang:golang@tcp(db:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB
func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}else {
		log.Println("Database connection established")
	}
}