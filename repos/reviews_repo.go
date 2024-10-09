package repos

import (
	"gorm.io/gorm"
	"movie-api/models"
)

type ReviewRepo struct {
	*Repository
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{&Repository{Db: db}}
}

func (r *ReviewRepo) List(obj *[]models.Review) {
	r.Db.Find(obj)
}
func (r *ReviewRepo) Get(id interface{}, obj *models.Review) {
	r.Db.Preload("User").Preload("Movie").Where("id = ?", id).Find(obj)
}
