package database

import (
	"log"
	"os"

	"github.com/verma29897/students_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB establishes a connection to the database and performs auto-migration
func ConnectDB() {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := "postgresql://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME

	if dsn == "postgresql://:@:/" {
		log.Fatal("DATABASE_URL environment variable is not set or is incomplete")
	}

	//log.Println("Database connection string:", dsn)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the Student model
	if err := DB.AutoMigrate(&models.Student{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection established and migration complete")
}

func CloseDB() {
	if db, err := DB.DB(); err == nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
