package database

import (
	"fmt"
	"os"
	"rentoday.id/app/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Start() {
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("NAME")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic("connection to db failed")
	} else {
		fmt.Println("success connect to db")
	}
	db.AutoMigrate(&database.Wardrobe{})
	fmt.Println("berhasil migrate")
}