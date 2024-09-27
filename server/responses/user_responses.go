package responses

import "movie-api/models"

type UserResponse struct {
	ID       uint                 `json:"id"`
	Name     string               `json:"name"`
	Username string               `json:"username"`
	Reviews  []UserReviewResponse `json:"reviews"`
}
type UserReviewResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Score   uint   `json:"score"`
	MovieID uint   `json:"movieID"`
}

func NewUserResponse(user *models.User) *UserResponse {
	userData := &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
	//userData.Reviews = make([]ReviewResponse, len(user.Reviews))
	//userData.Reviews = NewReviewsResponse(user.Reviews)
	for _, review := range user.Reviews {
		userData.Reviews = append(userData.Reviews, UserReviewResponse{
			ID:      review.ID,
			Title:   review.Title,
			Content: review.Content,
			Score:   review.Score,
			MovieID: review.MovieID,
		})
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
