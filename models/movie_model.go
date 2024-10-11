package models

type Movie struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	ImageURL    string `json:"image_url"`

	Reviews []Review
}
