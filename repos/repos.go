package repos

import "gorm.io/gorm"

type Repos struct {
	Db     *gorm.DB
	Movie  *MovieRepo
	Review *ReviewRepo
	User   *UserRepo
}

func NewRepos(db *gorm.DB) *Repos {
	return &Repos{
		Db:     db,
		Movie:  NewMovieRepo(db),
		Review: NewReviewRepo(db),
		User:   NewUserRepo(db),
	}
}
