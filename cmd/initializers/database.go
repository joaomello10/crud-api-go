package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var credentials = "host=localhost user=postgres password=1234 dbname=yourdb port=5432 sslmode=disable"

func ConnectToDB() *gorm.DB {
	dsn := credentials
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Successfully connected to database")
	return db
}
