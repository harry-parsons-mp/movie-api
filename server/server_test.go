package server_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"movie-api/handlers"
	"movie-api/models"
	"movie-api/repos"
	"movie-api/server"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	os.Setenv("TEST_MODE", "1")
	defer os.Unsetenv("TEST_MODE") // Clean up after test

	testServer := server.Server{}
	testServer.InitialiseDB("test_db.db")
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
	t.Run("add a movie", func(t *testing.T) {
		testMovie := models.Movie{
			Name:        "Test movie",
			Description: "Test desc",
			Genre:       "Testing",
		}
		jsonMovie, err := json.Marshal(testMovie)
		if err != nil {
			t.Fatalf("Failed to marshal movie: %v", err)
		}
		req := httptest.NewRequest(http.MethodPost, "/movie", strings.NewReader(string(jsonMovie)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := testServer.E.NewContext(req, rec)
		if err := movieHandler.CreateMovie(c); err != nil {
			t.Fatalf("Handler returned an error: %v", err)
		}
		got := &models.Movie{}
		err = json.NewDecoder(rec.Body).Decode(got)
		if err != nil {
			fmt.Errorf("failed to decode json %s", err.Error())
		}
		testMovieID := got.ID
		log.Println(got)
		if compareMovies(t, got, &testMovie) == false {
			t.Errorf("got %v, want %v", got, testMovie)
		}

		t.Run("Check if the movie just added can be fetched with a GET request", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := testServer.E.NewContext(req, rec)
			c.SetPath("/movie/:id")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%d", testMovieID))

			if err := movieHandler.GetMovieByID(c); err != nil {
				t.Fatalf("handler error: %v", err)
			}
			got := &models.Movie{}
			err = json.NewDecoder(rec.Body).Decode(got)
			if err != nil {
				t.Fatalf("failed to decode json %s", err.Error())
			}
			if compareMovies(t, got, &testMovie) == false {
				t.Errorf("got %v, want %v", got, testMovie)
			}
		})
	})
	os.Remove("test_db.db")
	testServer.E.Close()
}

func compareMovies(t *testing.T, got, want *models.Movie) bool {
	t.Helper()
	if got.Name == want.Name && got.Description == want.Description && got.Genre == want.Genre {
		return true
	}
	return false
}
