package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"movie-api/handlers"
	"movie-api/models"
	"movie-api/repos"
	"net/http"
	"os"
)

type Server struct {
	Db         *gorm.DB
	E          *echo.Echo
	MovieRepo  repos.MovieRepo
	ReviewRepo repos.ReviewRepo
}

func (server *Server) InitialiseDB(path string) {
	server.E = echo.New()

	Db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Db")
	}
	Db.AutoMigrate(&models.Movie{}, &models.Rating{}, &models.Review{}, &models.User{})
	server.Db = Db

}
func (server *Server) CloseDB() {
	db, err := server.Db.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Close()

}

func (server *Server) ConfigCors() {
	server.E.Use(middleware.Logger())
	server.E.Use(middleware.Recover())

	server.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

}

func (server *Server) InitialiseRoutes() {
	server.E.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	//Movies
	server.MovieRepo = repos.NewMovieRepo(server.Db)
	movieHandler := handlers.NewMovieHandler(server.MovieRepo)

	server.E.GET("/movies", movieHandler.GetAllMovies)
	server.E.GET("/movie/:id", movieHandler.GetMovieByID)
	server.E.POST("/movie", movieHandler.CreateMovie)
	server.E.PUT("/movie", movieHandler.UpdateMovie)
	server.E.DELETE("/movie", movieHandler.DeleteMovie)
	//// Reviews
	server.ReviewRepo = repos.NewReviewRepo(server.Db)
	reviewHandler := handlers.NewReviewHandler(server.ReviewRepo)

	server.E.GET("/reviews", reviewHandler.GetAllReviews)
	server.E.GET("/review/:id", reviewHandler.GetReviewByID)
	server.E.POST("review", reviewHandler.CreateReview)
	server.E.PUT("/review", reviewHandler.UpdateReview)
	server.E.DELETE("/review", reviewHandler.DeleteReview)

	//// Users
	//server.E.GET("/user/:username", Server.GetUser)
	//server.E.POST("/user", Server.AddUser)

	if os.Getenv("TEST_MODE") == "" {
		server.E.Logger.Fatal(server.E.Start(":1234"))
	}
}
