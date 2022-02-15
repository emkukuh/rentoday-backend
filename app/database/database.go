package database

import (
	"fmt"
	"os"
	"rentoday.id/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Start() {
	host := "rentoday-backend-rentoday-postgres-1"
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	dbURI := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, "password", dbPort)
	DB, err = gorm.Open(postgres.New(postgres.Config{DSN: dbURI}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(
		&model.Wardrobe{},
		&model.WardrobeCategory{},
		&model.WardrobeMaterial{},
		&model.User{},
	)
}