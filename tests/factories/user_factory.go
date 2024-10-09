package factories

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"movie-api/models"
)

func UserFactory(db *gorm.DB, u *models.User) *models.User {
	if u.Name == "" {
		u.Name = fmt.Sprintf("Name %d", rand.Int())
	}
	if u.Username == "" {
		u.Username = fmt.Sprintf("Username %d", rand.Int())
	}

	db.Create(&u)

	return u
}
