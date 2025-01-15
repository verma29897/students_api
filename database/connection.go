package database

import (
	"fmt"
	"log"
	"time"

	"github.com/verma29897/students-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dbConfig config.DatabaseConfig) {
	// Validate config
	if dbConfig.Host == "" || dbConfig.User == "" || dbConfig.Password == "" || dbConfig.Name == "" {
		log.Fatal("Database configuration is incomplete.")
	}

	// Build connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port, dbConfig.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Set connection pooling
	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	log.Println("Successfully connected to the database!")
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
