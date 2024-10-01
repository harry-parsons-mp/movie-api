package models

type Movie struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`

	Reviews []Review
}
