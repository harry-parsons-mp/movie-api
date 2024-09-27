package tests

import (
	"fmt"
	"movie-api/models"
	"net/http"
	"testing"
)

func TestReviewCreate(t *testing.T) {
	ts.ClearTable("reviews")

	request := Request{
		Method: http.MethodPost,
		Url:    "/reviews",
	}
	newReview := models.Review{}
	ReviewFactory(&newReview, 1, 1)

	// testing review factory without movieID or userID
	factorytest := models.Review{}
	ReviewFactory(&factorytest, 0, 0)

	tests := []TestCase{
		{
			TestName:    "Can create a new review",
			Request:     request,
			RequestBody: newReview,
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
		{
			TestName:    "Factory test without movieid or userid",
			Request:     request,
			RequestBody: factorytest,
			Expected: ExpectedResponse{
				StatusCode: http.StatusCreated,

				BodyParts: []string{
					fmt.Sprintf(`"title":"%v"`, factorytest.Title),
					fmt.Sprintf(`"content":"%v"`, factorytest.Content),
					fmt.Sprintf(`"score":%d`, factorytest.Score),
					fmt.Sprintf(`"userID":%d`, factorytest.UserID),
					fmt.Sprintf(`"movieID":%d`, factorytest.MovieID),
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
	review := models.Review{}
	ReviewFactory(&review, 1, 1)
	ts.S.Db.Create(&review)
	reviewID := review.ID
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
					fmt.Sprintf(`"userID":%d`, review.UserID),
					fmt.Sprintf(`"movieID":%d`, review.MovieID),
				},
			}},
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
					fmt.Sprintf(`"userID":%d`, review.UserID),
					fmt.Sprintf(`"movieID":%d`, review.MovieID),
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
func TestReviewUpdate(t *testing.T) {
	ts.ClearTable("reviews")
	//add a review to update
	review := models.Review{}
	ReviewFactory(&review, 1, 1)
	ts.S.Db.Create(&review)
	reviewID := review.ID
	request := Request{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/reviews/%d", +reviewID),
	}
	updatedReview := models.Review{
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
			RequestBody: updatedReview,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					updatedReview.Title,
					updatedReview.Content,
					fmt.Sprint(updatedReview.Score),
					fmt.Sprint(updatedReview.MovieID),
					fmt.Sprint(updatedReview.UserID)},
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
			RequestBody: updatedReview,
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
			RequestBody: updatedReview,
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
	ReviewFactory(&review, 1, 1)
	ts.S.Db.Create(&review)
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
			TestName: "Can't delete a review that isn't in the db",
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
			TestName: "Can't delete without an id",
			Request: Request{
				Method: http.MethodDelete,
				Url:    "/reviews",
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusMethodNotAllowed,
			},
		},
		{
			TestName: "Can't delete a review with a non integer id",
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
