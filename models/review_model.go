package models

type Review struct {
	ID      uint   `gorm:"primarykey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Score   uint   `json:"score"`
	UserID  uint   `json:"userID,omitempty"`
	MovieID uint   `json:"movieID,omitempty"`

	User  User  `gorm:"foreignkey:UserID" json:"user,omitempty"`
	Movie Movie `gorm:"foreignkey:MovieID" json:"movie,omitempty"`
}
