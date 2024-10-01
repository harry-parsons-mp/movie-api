package factories

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"movie-api/models"
)

func ReviewFactory(db *gorm.DB, r *models.Review) *models.Review {
	if r.Title == "" {
		r.Title = fmt.Sprintf("Review title %d", rand.Int())
	}
	if r.Content == "" {
		r.Content = fmt.Sprintf("Review description %d", rand.Int())
	}
	if r.Score == 0 {
		r.Score = uint(rand.Intn(11))
	}
	db.Create(&r)
	return r

}
