package tests

import (
	"fmt"
	"movie-api/models"
	"movie-api/server/requests"
	"movie-api/tests/factories"
	"net/http"
	"testing"
)

func TestUserList(t *testing.T) {
	ts.ClearTable("users")

	review := models.Review{}
	factories.ReviewFactory(ts.S.Db, &review)
	reviews := []models.Review{
		review,
	}
	// create user
	user := models.User{
		Reviews: reviews,
	}
	factories.UserFactory(ts.S.Db, &user)
	// create review

	request := Request{
		Method: http.MethodGet,
		Url:    "/users",
	}

	tests := []TestCase{
		{
			TestName: "Can list users",
			Request:  request,

			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, user.Name),
					fmt.Sprintf(`"username":"%v"`, user.Username),
					fmt.Sprintf(`"title":"%v"`, review.Title),
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

func TestUserGet(t *testing.T) {
	ts.ClearTable("users")
	ts.ClearTable("reviews")

	//create user and review:
	user := models.User{}
	factories.UserFactory(ts.S.Db, &user)

	review := models.Review{}
	factories.ReviewFactory(ts.S.Db, &review)

	id := user.ID
	tests := []TestCase{

		{
			TestName: "Can get user by id",
			Request: Request{
				Method: http.MethodGet,
				Url:    fmt.Sprintf("/users/%d", id),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, user.Name),
					fmt.Sprintf(`"username":"%v"`, user.Username),
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

func TestUserCreate(t *testing.T) {
	ts.ClearTable("users")
	request := Request{
		Method: http.MethodPost,
		Url:    "/users",
	}
	newUser := models.User{}
	factories.UserFactory(ts.S.Db, &newUser)
	newUserReq := requests.UserRequest{
		Name:     newUser.Name,
		Username: newUser.Username,
	}
	tests := []TestCase{
		{
			TestName:    "Can create user",
			Request:     request,
			RequestBody: newUserReq,
			Expected: ExpectedResponse{
				StatusCode: http.StatusCreated,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, newUser.Name),
					fmt.Sprintf(`"username":"%v"`, newUser.Username),
				},
			},
		},
		{
			TestName:    "Can't create user without username",
			Request:     request,
			RequestBody: models.User{Name: "testname"},
			Expected: ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "username is required",
			},
		},
		{
			TestName:    "Can't create with incorrect fields",
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

func TestUserUpdate(t *testing.T) {
	//clear the table
	ts.ClearTable("users")
	// add a user to be updated
	user := models.User{}
	factories.UserFactory(ts.S.Db, &user)
	id := user.ID

	request := Request{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/users/%d", id),
	}
	updatedUser := models.User{
		Name:     "Updated name",
		Username: "Updated username",
	}
	updatedUserReq := requests.UserRequest{
		Name:     updatedUser.Name,
		Username: updatedUser.Username,
	}
	tests := []TestCase{
		{
			TestName:    "Can update user",
			Request:     request,
			RequestBody: updatedUserReq,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%d`, id),
					fmt.Sprintf(`"name":"%v"`, updatedUser.Name),
					fmt.Sprintf(`"username":"%v"`, updatedUser.Username),
				},
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
			TestName: "Cannot update user that does not exist",
			Request: Request{
				Method: http.MethodPut,
				Url:    fmt.Sprintf("/users/%d", id+1),
			},
			RequestBody: updatedUser,
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("Failed to find user with id:%d", id+1),
			},
		},
		{
			TestName: "Cannot update a user with an invalid ID",
			Request: Request{
				Method: http.MethodPut,
				Url:    "/users/invalid-id",
			},
			RequestBody: updatedUser,
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Failed to find user with id:invalid-id",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			RunTest(t, test, ts)

		})
	}
}

func TestUserDelete(t *testing.T) {

	ts.ClearTable("users")
	//create a user to be deleted
	user := models.User{}
	factories.UserFactory(ts.S.Db, &user)
	id := user.ID

	request := Request{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/users/%v", id),
	}

	tests := []TestCase{
		{
			TestName: "Can delete user of an id that exists",
			Request:  request,
			Expected: ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyPart:   fmt.Sprintf("User of id: %d deleted sucessfully", id),
			},
		},
		{
			TestName: "Can't delete a user that isn't in the db",
			Request: Request{
				Method: http.MethodDelete,
				Url:    fmt.Sprintf("/users/%d", id+1),
			},
			Expected: ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   fmt.Sprintf("failed to find user of id: %v", id+1),
			},
		},
		{
			TestName: "Can't delete a movie with invalid id",
			Request: Request{
				Method: http.MethodDelete,
				Url:    "/users/test",
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
