package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = os.Getenv("DSN")
var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}
}
