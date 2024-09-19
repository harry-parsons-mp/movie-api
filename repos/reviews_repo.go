package repos

import (
	"gorm.io/gorm"
	"movie-api/models"
)

type reviewRepo struct {
	DB *gorm.DB
}
type ReviewRepo interface {
	GetAllReviews() ([]models.Review, error)
	GetReviewByID(id uint) (models.Review, error)
	CreateReview(review *models.Review) error
	UpdateReview(review *models.Review) error
	DeleteReview(review *models.Review) error
}

func NewReviewRepo(db *gorm.DB) ReviewRepo {
	return &reviewRepo{DB: db}
}

func (r *reviewRepo) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := r.DB.Preload("User").Preload("Movie").Find(&reviews).Error
	return reviews, err
}
func (r *reviewRepo) GetReviewByID(id uint) (models.Review, error) {
	var review models.Review
	err := r.DB.Preload("User").Preload("Review").Find(&review, "ID = ?", id).Error
	return review, err

}
func (r *reviewRepo) CreateReview(obj *models.Review) error {
	return r.DB.Create(&obj).Error
}
func (r *reviewRepo) UpdateReview(obj *models.Review) error {
	return r.DB.Updates(&obj).Error
}
func (r *reviewRepo) DeleteReview(obj *models.Review) error {
	return r.DB.Delete(obj).Error
}
