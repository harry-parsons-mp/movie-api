package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string
	Description string
	Genre       string
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
	Title   string
	Content string
	UserID  uint
	MovieID uint
}
type Rating struct {
	score   uint
	UserID  uint
	MovieID uint
}
