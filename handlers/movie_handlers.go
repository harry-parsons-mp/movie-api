package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"movie-api/models"
	"movie-api/repos"
	"movie-api/utils"
	"net/http"
)

type MovieHandler struct {
	Repo repos.MovieRepo
}
type GenericHandler[T any] struct {
	Repo repos.Repo[T]
}

func NewMovieHandler(repo repos.MovieRepo) *MovieHandler {
	return &MovieHandler{Repo: repo}
}

//func NewGenericHandler(repo repos.Repo[T]) *GenericHandler[T] {
//	return &GenericHandler[T]{Repo: repo}
//}

func (h *GenericHandler[T]) GetAllObj(c echo.Context) error {
	obj, err := h.Repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve movies")
	}
	return c.JSON(http.StatusOK, obj)
}

func (h *MovieHandler) GetAllMovies(c echo.Context) error {
	movies, err := h.Repo.GetAllMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve movies")
	}
	return c.JSON(http.StatusOK, movies)
}
func (h *MovieHandler) GetMovieByID(c echo.Context) error {
	id := utils.Str_to_uint(c.Param("id"))
	movie, err := h.Repo.GetMovieByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to retreive movie of id = %d", id))
	}
	return c.JSON(http.StatusOK, movie)

}
func (h *MovieHandler) CreateMovie(c echo.Context) error {
	var movie models.Movie

	name := c.FormValue("name")
	if name != "" {
		movie.Name = name
	}
	description := c.FormValue("description")
	if description != "" {
		movie.Description = description
	}
	genre := c.FormValue("genre")
	if genre != "" {
		movie.Genre = genre
	}
	err := h.Repo.CreateMovie(&movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add movie")
	}
	return c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) UpdateMovie(c echo.Context) error {
	var m models.Movie
	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := h.Repo.UpdateMovie(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update movie")
	}
	return c.JSON(http.StatusOK, m)
}

func (h *MovieHandler) DeleteMovie(c echo.Context) error {
	var m models.Movie
	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := h.Repo.DeleteMovie(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update movie")
	}
	return c.JSON(http.StatusOK, m)
}
