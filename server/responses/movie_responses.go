package responses

import (
	"movie-api/models"
)

type MovieResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Genre       string           `json:"genre"`
	ImageURL    string           `json:"image_url"`
	Review      []ReviewResponse `json:"reviews"`
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
		ImageURL:    movie.ImageURL,
	}
	if movie.Reviews != nil {
		res.Review = NewReviewsResponse(movie.Reviews)
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
