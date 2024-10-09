package factories

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"movie-api/models"
)

func MovieFactory(db *gorm.DB, m *models.Movie) *models.Movie {
	if m.Name == "" {
		m.Name = fmt.Sprintf("Movie name %d", rand.Int())
	}
	if m.Description == "" {
		m.Description = fmt.Sprintf("Movie description %d", rand.Int())
	}
	if m.Genre == "" {
		m.Genre = fmt.Sprintf("Movie genre %d", rand.Int())
	}

	db.Create(&m)

	return m
}
