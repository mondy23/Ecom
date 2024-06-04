package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBconn *gorm.DB
	err    error
)

func InitDB() {
	dsn := "host=localhost user=postgres password=fdsap@2024 dbname=w3schools port=5432 sslmode=disable timezone=asia/manila"
	DBconn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Can't connect to database")
	}
}
