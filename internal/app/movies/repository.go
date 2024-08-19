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
	DeleteMovie(ctx context.Context, movieId int) error
	UpdateMovie(ctx context.Context, movie Movie, movieId int) error
	GetMovies(ctx context.Context) ([]Movie, error)
	GetMovieDetailsById(ctx context.Context, movieId int) (*Movie, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindMovieByNameAndLanguage(ctx context.Context, movie Movie) (*Movie, error) {
	movieData := &Movie{}

	result := r.db.Where("title ILIKE ? AND language ILIKE ?", movie.Title, movie.Language).First(&movieData)
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
func (r *repository) DeleteMovie(ctx context.Context, movieId int) error {
	movie := Movie{}
	if err := r.db.Where("id=?", movieId).Delete(&movie).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetMovieDetailsById(ctx context.Context, movieId int) (*Movie, error) {
	movie := Movie{}
	if err := r.db.Where("id=?", movieId).First(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *repository) GetMovies(ctx context.Context) ([]Movie, error) {
	movies := []Movie{}
	res := r.db.Find(&movies)
	if res.Error != nil {
		return nil, res.Error
	}
	return movies, nil
}
func (r *repository) UpdateMovie(ctx context.Context, movie Movie, movieId int) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)

	result := r.db.Model(&Movie{}).Where("id = ?", movieId).Updates(movie)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
