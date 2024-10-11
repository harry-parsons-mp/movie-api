package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"movie-api/models"
	"movie-api/server/requests"
	"movie-api/tests/factories"
	"net/http"
	"testing"
)

func TestMovieList(t *testing.T) {
	ts.ClearTable("movies")

	movie := &models.Movie{}
	factories.MovieFactory(ts.S.Db, movie)

	tests := []TestCase{
		{
			TestName: "Can list all movies",
			Request: Request{
				Method: http.MethodGet,
				Url:    "/movies",
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, movie.Name),
					fmt.Sprintf(`"description":"%v"`, movie.Description),
					fmt.Sprintf(`"genre":"%v"`, movie.Genre),
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})
	}
}

func TestMovieGet(t *testing.T) {
	ts.ClearTable("movies")
	ts.ClearTable("reviews")

	// Create movie with review
	movie := &models.Movie{}
	factories.MovieFactory(ts.S.Db, movie)
	movieID := movie.ID

	// Create some reviews for the movie

	user1 := &models.User{}
	factories.UserFactory(ts.S.Db, user1)
	review1 := &models.Review{
		UserID:  user1.ID,
		MovieID: movieID,
	}

	factories.ReviewFactory(ts.S.Db, review1)
	review2 := &models.Review{
		UserID:  user1.ID,
		MovieID: movieID,
	}
	factories.ReviewFactory(ts.S.Db, review2)

	tests := []TestCase{
		{
			TestName: "Can get a movie by id",
			Request: Request{
				Method: http.MethodGet,
				Url:    fmt.Sprintf("/movies/%d", movieID),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, movie.Name),
					fmt.Sprintf(`"description":"%v"`, movie.Description),
					fmt.Sprintf(`"genre":"%v"`, movie.Genre),
					fmt.Sprintf(`"title":"%v"`, review1.Title),
					fmt.Sprintf(`"title":"%v"`, review2.Title),
					fmt.Sprintf(`"username":"%v"`, user1.Username),
				},
			},
		},
		{
			TestName: "Can get a movie that does not exist",
			Request: Request{
				Method: http.MethodGet,
				Url:    fmt.Sprintf("/movies/%d", movieID+1),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("Failed to retreive movie of id = %v", movieID+1),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})
	}
}

func TestMovieCreate(t *testing.T) {
	ts.ClearTable("movies")

	request := Request{
		Method: http.MethodPost,
		Url:    "/movies",
	}
	newMovie := &models.Movie{}
	factories.MovieFactory(ts.S.Db, newMovie)
	movieRequest := requests.MovieRequest{
		Data: requests.MovieRequestData{
			Name:        newMovie.Name,
			Description: newMovie.Description,
			Genre:       newMovie.Genre,
		},
	}
	tests := []TestCase{
		{
			TestName:    "can create movie",
			Request:     request,
			RequestBody: movieRequest,
			Expected: ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, newMovie.Name),
					fmt.Sprintf(`"description":"%v"`, newMovie.Description),
					fmt.Sprintf(`"genre":"%v"`, newMovie.Genre),
				},
				Callback: func(t *testing.T) {
					// DO Stuff here
					movieExists := &models.Movie{}
					ts.S.Db.Where("name = ? AND description = ? AND genre = ?", newMovie.Name, newMovie.Description, newMovie.Genre).Find(movieExists)
					// Ensure the movie actually exists in the DB
					assert.NotEqual(t, movieExists.ID, 0, "Expected movie to be found")
				},
			},
		},
		{
			TestName: "can't create movie without name",
			Request:  request,

			RequestBody: models.Movie{Description: "test dec", Genre: "test genre"},
			Expected: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "name of movie required",
			},
		},
		{
			TestName:    "can't create movie with incorrect fields",
			Request:     request,
			RequestBody: "hello world",
			Expected: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})

	}
}

func TestMovieUpdate(t *testing.T) {

	ts.ClearTable("movies")

	movie := &models.Movie{}
	factories.MovieFactory(ts.S.Db, movie)

	request := Request{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/movies/%v", movie.ID),
	}

	tests := []TestCase{
		{
			TestName:    "Cannot update a movie with incorrect fields",
			Request:     request,
			RequestBody: "incorrect fields",
			Expected: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "cannot unmarshal string",
			},
		},
		{
			TestName: "Cannot update a movie that does not exist",
			Request: Request{
				Method: http.MethodPut,
				Url:    "/movies/9999",
			},
			RequestBody: requests.MovieRequest{
				Data: requests.MovieRequestData{
					Name:        "Updated Title",
					Description: "Updated Desc",
					Genre:       "Updated Genre",
				},
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Failed to find movie with id",
			},
		},
		{
			TestName: "Cannot update a movie with an invalid ID",
			Request: Request{
				Method: http.MethodPut,
				Url:    "/movies/invalid-id",
			},
			RequestBody: requests.MovieRequest{
				Data: requests.MovieRequestData{
					Name:        "Updated Title",
					Description: "Updated Desc",
					Genre:       "Updated Genre",
				},
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Failed to find movie with id",
			},
		},
		{
			TestName: "Can update movie",
			Request:  request,
			RequestBody: requests.MovieRequest{
				Data: requests.MovieRequestData{
					Name:        "Updated Title",
					Description: "Updated Desc",
					Genre:       "Updated Genre",
				},
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%v`, movie.ID),
					`"name":"Updated Title"`,
					`"description":"Updated Desc"`,
					`"genre":"Updated Genre"`,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)
		})
	}
}

func TestMovieDelete(t *testing.T) {
	//clear the tables
	ts.ClearTable("movies")

	// create a movie to delete
	movie := &models.Movie{}
	factories.MovieFactory(ts.S.Db, movie)
	movieID := movie.ID
	request := Request{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/movies/%v", movieID),
	}

	tests := []TestCase{
		{
			TestName: "Can delete movie of an id that exists",
			Request:  request,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   fmt.Sprintf("Movie of id: %v deleted sucessfully", movieID),
			},
		},
		{
			TestName: "Can't delete a movie that does not exist",
			Request: Request{
				Method: http.MethodDelete,
				Url:    fmt.Sprintf("/movies/%d", movieID+1),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("failed to find movie of id: %v", movieID+1),
			},
		},
		{
			TestName: "Can't delete a movie with invalid id",
			Request: Request{
				Method: http.MethodDelete,
				Url:    "/movies/test",
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})

	}
}
