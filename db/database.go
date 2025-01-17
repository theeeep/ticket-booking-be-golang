package db

import (
	"fmt"
	"log"

	"github.com/theeeep/ticket-booking-be/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, dbMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=5432",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)

	log.Printf("Attempting to connect to database with URI: host=%s dbname=%s user=%s sslmode=%s",
		config.DBHost, config.DBName, config.DBUser, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Successfully connected to database: %s", config.DBName)

	log.Printf("Running database migrations...")
	if err := dbMigrator(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Printf("Database migrations completed successfully")

	return db
}
