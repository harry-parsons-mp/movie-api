package responses

import (
	"movie-api/models"
)

type MovieResponse struct {
	ID          uint                  `json:"ID"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Genre       string                `json:"genre"`
	Review      []MovieReviewResponse `json:"reviews"`
}

type MovieReviewResponse struct {
	ID      uint         `json:"id"`
	Title   string       `json:"title"`
	Content string       `json:"content"`
	Score   uint         `json:"score"`
	UserID  uint         `json:"userID"`
	User    UserResponse `json:"user"`
}

func NewMovieResponse(movie *models.Movie) *MovieResponse {
	res := &MovieResponse{
		ID:          movie.ID,
		Name:        movie.Name,
		Description: movie.Description,
		Genre:       movie.Genre,
	}
	//res.Review = make([]ReviewResponse, len(movie.Reviews))
	//res.Review = NewReviewsResponse(movie.Reviews)

	for _, review := range movie.Reviews {
		res.Review = append(res.Review, MovieReviewResponse{
			ID:      review.ID,
			Title:   review.Title,
			Content: review.Content,
			Score:   review.Score,
			UserID:  review.UserID,
			User:    *NewUserResponse(&review.User),
		})
	}
	return res
}

func NewMoviesResponse(movies []models.Movie) []MovieResponse {
	var movieData []MovieResponse

	for _, movie := range movies {
		movieData = append(movieData, *NewMovieResponse(&movie))
	}

	return movieData
}
