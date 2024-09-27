package repos

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	*Repository
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{&Repository{Db: db}}
}

func (u *UserRepo) List(obj interface{}) {
	u.Db.Preload("Reviews").Find(obj)
}
func (u *UserRepo) Get(id, obj interface{}) {
	u.Db.Preload("Reviews").Where("id = ?", id).Find(obj)
}
