package repos

import (
	"gorm.io/gorm"
)

type ReviewRepo struct {
	*Repository
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{&Repository{Db: db}}
}

func (r *ReviewRepo) List(obj interface{}) {
	r.Db.Preload("User").Preload("Movie").Find(obj)
}
func (r *ReviewRepo) Get(id, obj interface{}) {
	r.Db.Preload("User").Preload("Movie").Where("id = ?", id).Find(obj)
}
