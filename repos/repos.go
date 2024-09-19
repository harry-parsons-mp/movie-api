package repos

import "gorm.io/gorm"

type Repo[T any] struct {
	db *gorm.DB
}

func NewRepo[T any](db *gorm.DB) *Repo[T] {
	return &Repo[T]{db: db}
}

func (r *Repo[T]) Create(obj *T) error {
	return r.db.Create(obj).Error
}

func (r *Repo[T]) GetAll() (T, error) {
	var obj T
	err := r.db.Find(&obj).Error
	return obj, err
}
