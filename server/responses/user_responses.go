package responses

import "movie-api/models"

type UserResponse struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Username string           `json:"username"`
	Reviews  []ReviewResponse `json:"reviews"`
}
type UserReviewResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Score   uint   `json:"score"`
	MovieID uint   `json:"movie_id"`
}

func NewUserResponse(user *models.User) *UserResponse {
	userData := &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
	if user.Reviews != nil {
		userData.Reviews = NewReviewsResponse(user.Reviews)
	}

	return userData
}

func NewUsersResponse(users []models.User) []UserResponse {
	var usersData []UserResponse
	for _, user := range users {
		usersData = append(usersData, *NewUserResponse(&user))
	}
	return usersData
}
