package repos

import (
	"gorm.io/gorm"
	"movie-api/models"
)

type MovieRepo struct {
	*Repository
}

func NewMovieRepo(db *gorm.DB) *MovieRepo {
	return &MovieRepo{&Repository{Db: db}}
}

func (r *MovieRepo) List(obj *[]models.Movie) {
	r.Db.Find(obj)

}

func (r *MovieRepo) Get(id interface{}, obj *models.Movie) {
	r.Db.Preload("Reviews.User").Where("id = ?", id).Find(obj)

}
