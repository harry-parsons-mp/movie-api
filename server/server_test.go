package server_test

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"movie-api/handlers"
	"movie-api/models"
	"movie-api/repos"
	"movie-api/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServer(t *testing.T) {
	os.Setenv("TEST_MODE", "1")
	defer os.Unsetenv("TEST_MODE") // Clean up after test

	testServer := server.Server{}
	testServer.InitialiseDB("test_db.tb")
	defer testServer.CloseDB()
	log.Println("Init routes")
	testServer.InitialiseRoutes()
	testServer.MovieRepo = repos.NewMovieRepo(testServer.Db)
	movieHandler := handlers.NewMovieHandler(testServer.MovieRepo)

	t.Run("return 200 OK", func(t *testing.T) {
		log.Println("starting test")
		req := httptest.NewRequest(http.MethodGet, "/movies", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := testServer.E.NewContext(req, rec)

		log.Println("Calling GetAllMovies")

		if err := movieHandler.GetAllMovies(c); err != nil {
			t.Fatalf("handler error: %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("got %d, want %d", rec.Code, http.StatusOK)
		}

	})
	t.Run("add a movie and check it is in the db", func(t *testing.T) {
		testMovie := models.Movie{Name: "TestMovie", Description: "Test desc", Genre: "Testing"}
		jsonMovie, err := json.Marshal(testMovie)
		if err != nil {
			panic(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/movie", bytes.NewBuffer(jsonMovie))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := testServer.E.NewContext(req, rec)

		got := &models.Movie{}
		derr := json.NewDecoder(rec.Body).Decode(got)
		if derr != nil {
			panic(derr)
		}
		log.Println(got.Name)

		if err := movieHandler.GetMovieByID(c); err != nil {
			t.Fatalf("handler error: %v", err)
		}
	})

	testServer.E.Close()
}


