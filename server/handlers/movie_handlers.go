package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"movie-api/models"
	"movie-api/server"
	"movie-api/server/requests"
	"movie-api/server/responses"
	"net/http"
)

type MovieHandler struct {
	server *server.Server
}

func NewMovieHandler(server *server.Server) *MovieHandler {
	return &MovieHandler{
		server: server,
	}
}

func (h *MovieHandler) List(c echo.Context) error {
	var movie []models.Movie
	h.server.Repos.Movie.List(&movie)

	//response
	res := responses.NewMoviesResponse(movie)
	return c.JSON(http.StatusOK, res)
}

func (h *MovieHandler) Get(c echo.Context) error {
	id := c.Param("id")
	movie := &models.Movie{}

	h.server.Repos.Movie.Get(id, movie)
	if movie.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to retreive movie of id = %v", id))
	}

	res := responses.NewMovieResponse(movie)
	return c.JSON(http.StatusOK, res)

}

func (h *MovieHandler) Create(c echo.Context) error {
	var req requests.MovieRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, "name of movie required")
	}
	// create movie:
	mov := &models.Movie{
		Name:        req.Name,
		Description: req.Description,
		Genre:       req.Genre,
	}

	err := h.server.Repos.Movie.Create(mov)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add movie")
	}

	// response
	res := responses.NewMovieResponse(mov)
	return c.JSON(http.StatusCreated, res)
}

func (h *MovieHandler) Update(c echo.Context) error {
	ID := c.Param("id")
	var updateRequest requests.MovieRequest

	if err := c.Bind(&updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check movie exists
	mov := &models.Movie{}
	h.server.Repos.Movie.Get(ID, mov)
	if mov.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to find movie with id: %d", ID))
	}

	// Update the movie

	mov.Name = updateRequest.Name
	mov.Description = updateRequest.Description
	mov.Genre = updateRequest.Genre

	err := h.server.Repos.Movie.Update(mov)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to update movie, %v", err))
	}

	// Return response
	res := responses.NewMovieResponse(mov)
	return c.JSON(http.StatusOK, res)
}

func (h *MovieHandler) Delete(c echo.Context) error {
	ID := c.Param("id")
	var toDelete models.Movie

	// check if movie exists
	h.server.Repos.Movie.Get(ID, &toDelete)
	if toDelete.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("failed to find movie of id: %v", ID))
	}

	err := h.server.Repos.Movie.Delete(&toDelete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update movie")
	}
	// create response
	res := fmt.Sprintf("Movie of id: %v deleted sucessfully", ID)
	return c.JSON(http.StatusOK, res)
}
