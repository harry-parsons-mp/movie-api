package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Reviews     []Review
	Ratings     []Rating
}

type User struct {
	gorm.Model
	Name     string
	Username string
	Reviews  []Review
	Ratings  []Rating
}

type Review struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
	User    User `gorm:"foreignkey:UserID"`
	MovieID uint
	Movie   Movie
}
type Rating struct {
	gorm.Model
	Score   uint
	UserID  uint
	User    User `json:"-"`
	MovieID uint
	Movie   Movie `json:"-"`
}
