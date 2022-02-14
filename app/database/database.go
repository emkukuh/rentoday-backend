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
	// host := os.Getenv("DB_HOST")
	host := "rentoday-backend-rentoday-postgres-1"
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, "password", dbPort)
	db, err = gorm.Open(postgres.New(postgres.Config{DSN: dbURI}), &gorm.Config{})
	if err != nil {
		panic("connection to db failed")
	} else {
		fmt.Println("success connect to db")
	}
	db.AutoMigrate(
		&database.Wardrobe{},
		&database.WardrobeCategory{},
		&database.WardrobeMaterial{},
	)
}