package server_test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"movie-api/models"
	"movie-api/server"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var db *gorm.DB

func TestServer(t *testing.T) {

	t.Run("return 200 OK", func(t *testing.T) {
		//e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		//c := e.NewContext(req, rec)

		if rec.Code != http.StatusOK {
			t.Errorf("got %d, want %d", rec.Code, http.StatusOK)
		}

	})
	t.Run("/movies returns the correct list of movies", func(t *testing.T) {
		e := echo.New()
		db = server.ConnectDB()

		req := httptest.NewRequest(http.MethodGet, "/movies", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		fmt.Println(rec)
		c := e.NewContext(req, rec)
		ListMovies(c)
		if rec.Code != http.StatusOK {
			t.Errorf("error connecting to the server")
		}
		var wantedMovies models.Movie
		db.Find(&wantedMovies)

		want := wantedMovies
		got := rec.Body.String()

		fmt.Println(got)
		if !reflect.DeepEqual(got, (want)) {
			t.Errorf("got: %+v want: %+v", got, want)
		}

	})

}
