package models

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Reviews []Review
}
