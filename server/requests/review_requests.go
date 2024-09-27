package requests

type ReviewRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Score   uint   `json:"score"`
	UserID  uint   `json:"userID"`
	MovieID uint   `json:"movieID"`
}

//type UpdateReviewRequest struct {
//	ID      uint   `json:"id"`
//	Title   string `json:"title"`
//	Content string `json:"content"`
//	Score   uint   `json:"score"`
//	UserID  uint   `json:"userID"`
//	MovieID uint   `json:"movieID"`
//}
