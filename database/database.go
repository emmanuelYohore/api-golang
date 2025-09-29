package database

import (
	"fmt"
	"log"

	"github.com/emmanuelYohore/api-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=api-golang port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connection established")
	DB.AutoMigrate(&models.Book{})

}
