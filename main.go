package main

import (
	"fmt"
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
	db.AutoMigrate(&models.Movie{}, &models.Rating{}, &models.Review{}, &models.User{})
	//server.MigrateDB(db, &models.Movie{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	e.GET("/movies", ListMovies)
	e.GET("/movie/:id", getMovieFromID)
	e.Logger.Fatal(e.Start(":1234"))
}

func ListMovies(c echo.Context) error {
	name := c.QueryParam("name")
	var movies []models.Movie
	//var reviews []models.Review
	if len(name) != 0 {
		db.Find(&movies, "Name = ?", name)
	} else {
		err := db.Model(&models.Movie{}).Preload("Users").Find(movies).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.JSON(http.StatusOK, movies)

}

func getMovieFromID(c echo.Context) error {
	id := c.Param("id")
	var movies []models.Movie
	db.First(&movies, "ID = ?", id)
	return c.JSON(http.StatusOK, movies)
}
