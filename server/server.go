package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"movie-api/handlers"
	"movie-api/models"
	"movie-api/repos"
	"net/http"
)

type Server struct {
	db         *gorm.DB
	e          *echo.Echo
	MovieRepo  repos.MovieRepo
	ReviewRepo repos.ReviewRepo
}

func MigrateDB[T any](database *gorm.DB, model T) {
	err := database.AutoMigrate(model)
	if err != nil {
		fmt.Errorf("failed to migrate")
	}

}

func (server *Server) InitialiseDB() {
	server.e = echo.New()

	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	db.AutoMigrate(&models.Movie{}, &models.Rating{}, &models.Review{}, &models.User{})
	server.db = db

}

func (server *Server) InitialiseRoutes() {
	server.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	//Movies
	server.MovieRepo = repos.NewMovieRepo(server.db)
	movieHandler := handlers.NewMovieHandler(server.MovieRepo)

	server.e.GET("/movies", movieHandler.GetAllMovies)
	server.e.GET("/movie/:id", movieHandler.GetMovieByID)
	server.e.POST("/movie", movieHandler.CreateMovie)
	server.e.PUT("/movie", movieHandler.UpdateMovie)
	server.e.DELETE("/movie", movieHandler.DeleteMovie)
	//// Reviews
	server.ReviewRepo = repos.NewReviewRepo(server.db)
	reviewHandler := handlers.NewReviewHandler(server.ReviewRepo)

	server.e.GET("/reviews", reviewHandler.GetAllReviews)
	server.e.GET("/review/:id", reviewHandler.GetReviewByID)
	server.e.POST("review", reviewHandler.CreateReview)
	server.e.PUT("/review", reviewHandler.UpdateReview)
	server.e.DELETE("/review", reviewHandler.DeleteReview)

	//// Users
	//server.e.GET("/user/:username", Server.GetUser)
	//server.e.POST("/user", Server.AddUser)

	server.e.Logger.Fatal(server.e.Start(":1234"))
}
