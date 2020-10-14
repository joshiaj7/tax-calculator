package config

import (
	"github.com/joshiaj7/tax-calculator/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB for global use
var DB *gorm.DB

// SetupDB to setup db
func SetupDB() (err error) {
	dsn := "host=localhost user=postgres password=password123 dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// create table
	DB.AutoMigrate(&model.Tax{})

	return err
}