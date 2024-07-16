package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = InitializePostgresSQL()
	if err != nil {
		return fmt.Errorf("Error initializing postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}
