package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"movie-api/models"
	"movie-api/server"
	"net/http"
)

var db *gorm.DB

func main() {
	e := echo.New()

	db = server.ConnectDB()
	server.MigrateDB(db, &models.Movie{})

	db.Create(&models.Movie{Name: "Movie 1", Description: "Movie 1 description", Genre: "Sci-Fi"})
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
