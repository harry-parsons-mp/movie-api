package tests

import (
	"fmt"
	"math/rand"
	"movie-api/models"
)

func MovieFactory(m *models.Movie) {
	if m.Name == "" {
		m.Name = fmt.Sprintf("Movie name %d", rand.Int())
	}
	if m.Description == "" {
		m.Description = fmt.Sprintf("Movie description %d", rand.Int())
	}
	if m.Genre == "" {
		m.Genre = fmt.Sprintf("Movie genre %d", rand.Int())
	}

}

func ReviewFactory(r *models.Review, movieID, userID uint) {
	if r.Title == "" {
		r.Title = fmt.Sprintf("Review title %d", rand.Int())
	}
	if r.Content == "" {
		r.Content = fmt.Sprintf("Review description %d", rand.Int())
	}
	if r.Score == 0 {
		r.Score = uint(rand.Intn(11))
	}
	if r.UserID == 0 {
		if userID != 0 {
			r.UserID = userID
		} else {
			tempUser := models.User{}
			UserFactory(&tempUser)
			r.UserID = tempUser.ID
		}

	}
	if r.MovieID == 0 {
		if movieID != 0 {
			r.MovieID = movieID
		} else {
			tempMovie := models.Movie{}
			MovieFactory(&tempMovie)
			r.MovieID = tempMovie.ID
		}

	}
}

func UserFactory(u *models.User) {
	if u.Name == "" {
		u.Name = fmt.Sprintf("Name %d", rand.Int())
	}
	if u.Username == "" {
		u.Username = fmt.Sprintf("Username %d", rand.Int())
	}
}
