package responses

import (
	"movie-api/models"
)

type ReviewResponse struct {
	ID      uint          `json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Score   uint          `json:"score"`
	UserID  uint          `json:"user_id,omitempty"`
	MovieID uint          `json:"movie_id,omitempty"`
	User    UserResponse  `json:"user,omitempty"`
	Movie   MovieResponse `json:"movie,omitempty"`
}

func NewReviewResponse(review *models.Review) *ReviewResponse {

	res := &ReviewResponse{
		ID:      review.ID,
		Title:   review.Title,
		Content: review.Content,
		Score:   review.Score,
	}
	if review.User.ID != 0 {
		res.User = *NewUserResponse(&review.User)
		res.UserID = review.UserID
	}
	if review.Movie.ID != 0 {
		res.Movie = *NewMovieResponse(&review.Movie)
		res.MovieID = review.MovieID
	}

	return res
}

func NewReviewsResponse(reviews []models.Review) []ReviewResponse {
	var reviewData []ReviewResponse

	for _, review := range reviews {
		reviewData = append(reviewData, *NewReviewResponse(&review))
	}

	return reviewData
}
