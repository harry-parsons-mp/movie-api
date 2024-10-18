package repos

import (
	"gorm.io/gorm"
	"movie-api/models"
)

type UserRepo struct {
	*Repository
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{&Repository{Db: db}}
}

func (u *UserRepo) List(obj *[]models.User) {
	u.Db.Preload("Reviews").Find(obj)
}

func (u *UserRepo) Get(id interface{}, obj *models.User) {
	u.Db.Preload("Reviews").Where("id = ?", id).Find(obj)
}

func (u *UserRepo) Auth(username interface{}, obj *models.User) {
	u.Db.Where("username = ?", username).Find(obj)
}
