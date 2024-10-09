package repos

import (
	"gorm.io/gorm"
	"movie-api/models"
)

type movieRepo struct {
	DB *gorm.DB
}
type MovieRepo interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovieByID(id uint) (models.Movie, error)
	CreateMovie(movie *models.Movie) error
	UpdateMovie(movie *models.Movie) error
	DeleteMovie(movie *models.Movie) error
}

func NewMovieRepo(db *gorm.DB) MovieRepo {
	return &movieRepo{DB: db}
}
func (r *movieRepo) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	err := r.DB.Preload("Reviews").Preload("Reviews.User").Preload("Ratings").Find(&movies).Error
	return movies, err
}

func (r *movieRepo) GetMovieByID(id uint) (models.Movie, error) {
	var movie models.Movie
	err := r.DB.Preload("Reviews").Preload("Reviews.User").Preload("Ratings").Find(&movie, "ID = ?", id).Error
	return movie, err

}
func (r *movieRepo) CreateMovie(m *models.Movie) error {
	return r.DB.Create(&m).Error
}
func (r *movieRepo) UpdateMovie(m *models.Movie) error {
	return r.DB.Updates(m).Error
}
func (r *movieRepo) DeleteMovie(m *models.Movie) error {
	return r.DB.Delete(m).Error
}
