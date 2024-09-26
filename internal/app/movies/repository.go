package movies

import (
	"context"
	"fmt"

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
	GetMoviesByLanguage(ctx context.Context, language string) ([]Movie, error)
	GetMoviesByGenre(ctx context.Context, genre string) ([]Movie, error)
	GetMovieByName(ctx context.Context, name string) (*Movie, error)
	GetMovieByNameAndLanguage(ctx context.Context, name, language string) (*Movie, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetMovieByNameAndLanguage(ctx context.Context, name, language string) (*Movie, error) {
	movieData := &Movie{}

	result := r.db.Where("title ILIKE ? AND language ILIKE ?", name, language).First(&movieData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movieData, nil
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
	result := r.db.Where("id=?", movieId).Delete(&movie)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no movie found with ID %d", movieId)
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
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *repository) GetMovieByName(ctx context.Context, name string) (*Movie, error) {
	movie := Movie{}
	result := r.db.Where("title = ?", name).First(&movie)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("no movie found with name %s", name)
	}
	return &movie, nil
}

func (r *repository) GetMoviesByGenre(ctx context.Context, genre string) ([]Movie, error) {
	movies := []Movie{}
	result := r.db.Where("genre ILIKE ?", genre).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}

func (r *repository) GetMoviesByLanguage(ctx context.Context, language string) ([]Movie, error) {
	movies := []Movie{}
	result := r.db.Where("language ILIKE ?", language).Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}
	return movies, nil
}
