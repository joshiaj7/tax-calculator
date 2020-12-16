package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/joshiaj7/tax-calculator/internal/model"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	hostname string `envconfig:"HOSTNAME"`
	user     string `envconfig:"USER"`
	password string `envconfig:"PASSWORD"`
	dbName   string `envconfig:"NAME"`
	port     string `envconfig:"PORT"`
}

// DB for global use
var DB *gorm.DB

// SetupDB to setup db
func SetupDB() (err error) {
	godotenv.Load(".env")

	var c config
	err = envconfig.Process("DB", &c)

	format := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	dsn := fmt.Sprintf(format, c.hostname, c.user, c.password, c.dbName, c.port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// create table
	DB.AutoMigrate(&model.Tax{})

	return err
}
