package repos

import (
	"gorm.io/gorm"
)

type MovieRepo struct {
	*Repository
}

func NewMovieRepo(db *gorm.DB) *MovieRepo {
	return &MovieRepo{&Repository{Db: db}}
}

func (r *MovieRepo) List(obj interface{}) {
	r.Db.Preload("Reviews").Preload("Reviews.User").Find(obj)

}

func (r *MovieRepo) Get(id, obj interface{}) {
	r.Db.Preload("Reviews").Where("id = ?", id).Find(obj)

}
