package tests

import (
	"fmt"
	"movie-api/models"
	"movie-api/server/requests"
	"movie-api/tests/factories"
	"net/http"
	"testing"
)

func TestReviewList(t *testing.T) {
	ts.ClearTable("reviews")
	review := models.Review{}
	factories.ReviewFactory(ts.S.Db, &review)

	request := Request{
		Method: http.MethodGet,
		Url:    "/reviews",
	}

	tests := []TestCase{
		{
			TestName: "Can list reviews",
			Request:  request,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"title":"%v"`, review.Title),
					fmt.Sprintf(`"content":"%v"`, review.Content),
					fmt.Sprintf(`"score":%d`, review.Score),
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

func TestReviewGet(t *testing.T) {
	// clear tables
	ts.ClearTable("reviews")
	// create a review to add and then fetch
	movie := models.Movie{}
	factories.MovieFactory(ts.S.Db, &movie)

	user := models.User{}
	factories.UserFactory(ts.S.Db, &user)

	review := models.Review{
		UserID:  user.ID,
		User:    user,
		MovieID: movie.ID,
		Movie:   movie,
	}
	factories.ReviewFactory(ts.S.Db, &review)
	reviewID := review.ID

	tests := []TestCase{
		{
			TestName: "Can get review",
			Request: Request{
				Method: http.MethodGet,
				Url:    fmt.Sprintf("/reviews/%d", reviewID),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%d`, review.ID),
					fmt.Sprintf(`"title":"%v"`, review.Title),
					fmt.Sprintf(`"content":"%v"`, review.Content),
					fmt.Sprintf(`"score":%d`, review.Score),
					fmt.Sprintf(`"user_id":%d`, user.ID),
					fmt.Sprintf(`"movie_id":%d`, movie.ID),
				},
			},
		},
		{
			TestName: "Cannot get review that does not exist",
			Request: Request{
				Method: http.MethodGet,
				Url:    fmt.Sprintf("/reviews/%d", reviewID+1),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("Failed to retreive review of id: %d", reviewID+1),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})

	}

}
func TestReviewCreate(t *testing.T) {
	ts.ClearTable("reviews")

	request := Request{
		Method: http.MethodPost,
		Url:    "/reviews",
	}
	newReview := models.Review{}
	factories.ReviewFactory(ts.S.Db, &newReview)
	reviewRequest := requests.ReviewRequest{
		Title:   newReview.Title,
		Content: newReview.Content,
		Score:   newReview.Score,
		UserID:  newReview.UserID,
		MovieID: newReview.MovieID,
	}

	tests := []TestCase{
		{
			TestName:    "Can create a new review",
			Request:     request,
			RequestBody: reviewRequest,
			Expected: ExpectedResponse{
				StatusCode: 201,
				BodyParts: []string{
					newReview.Title,
					newReview.Content,
					fmt.Sprintf("%d", newReview.Score),
					fmt.Sprintf("%d", newReview.UserID),
					fmt.Sprintf("%d", newReview.MovieID)},
			},
		},
		{
			TestName: "Can't create a review without a title",
			Request:  request,
			RequestBody: models.Review{
				Content: "Test",
				Score:   1,
				UserID:  1,
				MovieID: 1,
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "review title is required",
			},
		},
		{
			TestName:    "Can't create a review with wrong fields",
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

func TestReviewUpdate(t *testing.T) {
	ts.ClearTable("reviews")
	//add a review to update
	review := models.Review{}
	factories.ReviewFactory(ts.S.Db, &review)
	reviewID := review.ID

	request := Request{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/reviews/%d", +reviewID),
	}
	// creating a new request
	UpdatedReviewReq := requests.ReviewRequest{
		Title:   "Updated title",
		Content: "Updated content",
		Score:   2,
		UserID:  2,
		MovieID: 2,
	}
	tests := []TestCase{
		{
			TestName:    "Can update review",
			Request:     request,
			RequestBody: UpdatedReviewReq,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					UpdatedReviewReq.Title,
					UpdatedReviewReq.Content,
					fmt.Sprint(UpdatedReviewReq.Score),
					fmt.Sprint(UpdatedReviewReq.MovieID),
					fmt.Sprint(UpdatedReviewReq.UserID)},
			},
		},
		{
			TestName:    "Can't update with incorrect field",
			Request:     request,
			RequestBody: "incorrect fields",
			Expected: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "cannot unmarshal string",
			},
		},
		{
			TestName: "Cannot update movie that does not exist",
			Request: Request{
				Method: http.MethodPut,
				Url:    "/reviews/99999",
			},
			RequestBody: UpdatedReviewReq,
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("Failed to find review with id:99999"),
			},
		},
		{
			TestName: "Cannot update a review with an invalid ID",
			Request: Request{
				Method: http.MethodPut,
				Url:    "/reviews/invalid-id",
			},
			RequestBody: UpdatedReviewReq,
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Failed to find review with id:invalid-id",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})
	}
}

func TestReviewDelete(t *testing.T) {
	//clear table
	ts.ClearTable("reviews")
	review := models.Review{}
	factories.ReviewFactory(ts.S.Db, &review)
	id := review.ID

	request := Request{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/reviews/%v", id),
	}

	tests := []TestCase{
		{
			TestName: "Can delete review of an id that exists",
			Request:  request,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   fmt.Sprintf("Review of id: %d deleted sucessfully", id),
			},
		},
		{
			TestName: "Can't delete a review that doesn't exist",
			Request: Request{
				Method: http.MethodDelete,
				Url:    fmt.Sprintf("/reviews/%d", id+1),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("failed to find review of id: %v", id+1),
			},
		},
		{
			TestName: "Can't delete a review with invalid id",
			Request: Request{
				Method: http.MethodDelete,
				Url:    "/reviews/test",
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
