package config

import (
	"event-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("Environment variable 'DATABASE_URL' is not set.")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = database.AutoMigrate(&models.Event{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
	log.Println("Database connection established and migrated successfully.")
}
