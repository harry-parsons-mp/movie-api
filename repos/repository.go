package repos

import (
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) Create(obj interface{}) error {
	return r.Db.Create(obj).Error
}

func (r *Repository) List(obj interface{}) {
	r.Db.Find(obj)

}
func (r *Repository) Get(id, obj interface{}) {
	r.Db.Where("id = ?", id).Find(obj)
}

func (r *Repository) Update(object interface{}) error {
	return r.Db.Save(object).Error
}

func (r *Repository) Delete(obj interface{}) error {
	return r.Db.Delete(obj).Error
}
