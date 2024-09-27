package repos

import (
	"gorm.io/gorm"
	"log"
)

type MovieRepo struct {
	*Repository
}

func NewMovieRepo(db *gorm.DB) *MovieRepo {
	return &MovieRepo{&Repository{Db: db}}
}

func (r *MovieRepo) List(obj interface{}) {
	log.Println("In the movie repo!")
	r.Db.Preload("Reviews").Preload("Reviews.User").Find(obj)

}

func (r *MovieRepo) Get(id, obj interface{}) {
	log.Println("In the movie repo!")
	r.Db.Preload("Reviews").Where("id = ?", id).Find(obj)

}
