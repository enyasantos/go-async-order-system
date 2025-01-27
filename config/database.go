package config

import (
	"fmt"
	"log"
	"os"

	"github.com/enyasantos/go-async-order-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB instance: %v", err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		log.Fatalf("Postgres automigration error: %v", err)
	}

	return db, nil
}
