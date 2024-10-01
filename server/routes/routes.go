package routes

import (
	"github.com/labstack/echo/v4"
	"movie-api/server"
	"movie-api/server/handlers"
	"net/http"
)

const moviePath = "/movies"
const reviewPath = "/reviews"
const userPath = "/users"
const ratingPath = "/ratings"

func InitialiseRoutes(server *server.Server) {
	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	movieHandler := handlers.NewMovieHandler(server)
	reviewHandler := handlers.NewReviewHandler(server)
	userHandler := handlers.NewUserHandler(server)

	//Movies
	server.Echo.GET(moviePath, movieHandler.List)
	server.Echo.GET(moviePath+"/:id", movieHandler.Get)
	server.Echo.POST(moviePath, movieHandler.Create)
	server.Echo.PUT(moviePath+"/:id", movieHandler.Update)
	server.Echo.DELETE(moviePath+"/:id", movieHandler.Delete)

	// Reviews
	server.Echo.GET(reviewPath, reviewHandler.List)
	server.Echo.GET(reviewPath+"/:id", reviewHandler.GetByID)
	server.Echo.POST(reviewPath, reviewHandler.Create)
	server.Echo.PUT(reviewPath+"/:id", reviewHandler.Update)
	server.Echo.DELETE(reviewPath+"/:id", reviewHandler.Delete)

	// Users
	server.Echo.GET(userPath, userHandler.List)
	server.Echo.GET(userPath+"/:id", userHandler.Get)
	server.Echo.POST(userPath, userHandler.Create)
	server.Echo.PUT(userPath+"/:id", userHandler.Update)
	server.Echo.DELETE(userPath+"/:id", userHandler.Delete)

}
