package movies

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindMovieByNameAndLanguage(ctx context.Context, movie Movie) (*Movie, error)
	CreateMovie(ctx context.Context, movie Movie) (int, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindMovieByNameAndLanguage(ctx context.Context, movie Movie) (*Movie, error) {
	movieData := &Movie{}

	result := r.db.Where("title=? AND language=?", movie.Title, movie.Language).First(&movieData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movieData, nil
}

func (r *repository) CreateMovie(ctx context.Context, movie Movie) (int, error) {
	if err := r.db.Create(&movie).Error; err != nil {
		return 0, err
	}
	return int(movie.ID), nil
}
