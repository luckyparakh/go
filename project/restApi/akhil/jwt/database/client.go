package database

import (
	"jwt/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbErr error

func Connect(connString string) {
	Instance, dbErr = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
		panic("Error while connecting DB")
	}
	log.Print("Connection Sucessful")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Print("Migration complete")
}
