package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB akan membuat koneksi ke database SQLite dan menginisialisasi Gorm
func InitDB() (*gorm.DB, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Get database credential from .env file
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(sqlite.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}
