package responses

import "movie-api/models"

type ReviewResponse struct {
	ID      uint          `json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Score   uint          `json:"score"`
	UserID  uint          `json:"userID"`
	MovieID uint          `json:"movieID"`
	User    UserResponse  `json:"user"`
	Movie   MovieResponse `json:"movie"`
}

func NewReviewResponse(review *models.Review) *ReviewResponse {
	res := &ReviewResponse{
		ID:      review.ID,
		Title:   review.Title,
		Content: review.Content,
		Score:   review.Score,
		UserID:  review.UserID,
		MovieID: review.MovieID,
	}
	res.User = *NewUserResponse(&review.User)
	res.Movie = *NewMovieResponse(&review.Movie)
	return res
}

func NewReviewsResponse(reviews []models.Review) []ReviewResponse {
	var reviewData []ReviewResponse

	for _, review := range reviews {
		reviewData = append(reviewData, *NewReviewResponse(&review))
	}

	return reviewData
}
