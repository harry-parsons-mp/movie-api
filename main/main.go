package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"movie-api/models"
	"net/http"
)

var db *gorm.DB

func main() {
	e := echo.New()

	db = connectDB()
	MigrateDB(db, &models.Movie{})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	e.GET("/movies", ListMovies)
	e.Logger.Fatal(e.Start(":1234"))
}

func ListMovies(c echo.Context) error {
	var movies models.Movie
	db.Find(&movies)
	return c.JSON(http.StatusOK, movies)

}

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
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
