package main

import (
	"gorm/cmd/initializers"
	"gorm/cmd/models"
)

func main() {
	db := initializers.ConnectToDB()

	db.AutoMigrate(models.Note{})
}
