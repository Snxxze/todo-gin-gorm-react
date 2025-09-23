package configs

import (
	"backend/entity"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db = database
	log.Println("Database connected")
}

func SetupDatabase() {
	// Migrate
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Todo{},
	); err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}