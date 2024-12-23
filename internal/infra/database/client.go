package database

import (
	"backend/internal/domain/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Client() (*gorm.DB, error) {
	// load .env file
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error connecting to DB")
		return nil, nil
	}

	return db, nil
}
