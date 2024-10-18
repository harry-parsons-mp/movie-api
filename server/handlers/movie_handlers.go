package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"mime/multipart"
	"movie-api/models"
	"movie-api/server"
	"movie-api/server/requests"
	"movie-api/server/responses"
	"net/http"
	"os"
	"path/filepath"
)

var supportedTypes = map[string]bool{
	"image/jpeg": true,
	"image/jpg":  true,
	"image/png":  true,
}

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
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to retreive movie of id = %s", id))
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
		return echo.NewHTTPError(http.StatusBadRequest, "name of movie required")

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
	var req requests.MovieRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check movie exists
	mov := &models.Movie{}
	h.server.Repos.Movie.Get(ID, mov)
	if mov.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to find movie with id: %s", ID))
	}

	// Update the movie
	mov.Name = req.Name
	mov.Description = req.Description
	mov.Genre = req.Genre

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
	h.server.Repos.Movie.Get(ID, &toDelete)
	if toDelete.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("failed to find movie of id: %v", ID))
	}
	if toDelete.ImageURL != "" {
		os.Remove(toDelete.ImageURL)
	}

	err := h.server.Repos.Movie.Delete(&toDelete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update movie")
	}

	//response
	res := fmt.Sprintf("Movie of id: %v deleted sucessfully", ID)
	return c.JSON(http.StatusOK, res)
}

func (h *MovieHandler) ImageUpload(c echo.Context) error {
	id := c.FormValue("movieID")

	movie := &models.Movie{}
	h.server.Repos.Movie.Get(id, movie)

	if movie.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to find movie with id %v", id))
	}

	if movie.ImageURL != "" {
		log.Printf("deleting image %v ", movie.ImageURL)
		os.Remove(movie.ImageURL)

	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	path := "images/movie_" + fmt.Sprintf("%d", movie.ID) + filepath.Ext(file.Filename)

	err = handleImageUpload(file, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	movie.ImageURL = path

	err = h.server.Repos.Movie.Update(&movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to update movie, %v", err))
	}
	res := responses.NewMovieResponse(movie)
	return c.JSON(http.StatusOK, res)

}

func handleImageUpload(file *multipart.FileHeader, path string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//check the file type
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)

	if err != nil {
		return err
	}
	src.Seek(0, 0)
	contentType := http.DetectContentType(buffer)
	if _, ok := supportedTypes[contentType]; !ok {
		return errors.New("unsupported file type")
	}

	// Create destination file
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy file contents
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
