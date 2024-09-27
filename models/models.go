package models

type Movie struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`

	Reviews []Review
}

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Reviews []Review
}

type Review struct {
	ID      uint   `gorm:"primarykey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Score   uint   `json:"score"`
	UserID  uint   `json:"userID"`
	MovieID uint   `json:"movieID"`

	User  User  `gorm:"foreignkey:UserID"`
	Movie Movie `gorm:"foreignkey:MovieID"`
}
