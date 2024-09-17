package server

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	return db

}

func MigrateDB[T any](database *gorm.DB, model T) {
	err := database.AutoMigrate(model)
	if err != nil {
		fmt.Errorf("failed to migrate")
	}

}
