package movies

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type service struct {
	repo Repository
}

type Service interface {
	RegisterMovie(ctx context.Context, movie Movie) (int, error)
	UpdateMovie(ctx context.Context, movie Movie, movieId int) error
	ListMovies(ctx context.Context) ([]Movie, error)
	GetMovieDetailsByID(ctx context.Context, movieId int) (*Movie, error)
	DeleteMovie(ctx context.Context, movieId int) error
	GetMoviesByLanguage(ctx context.Context, language string) ([]Movie, error)
	GetMoviesByGenre(ctx context.Context, genre string) ([]Movie, error)
	GetMovieByName(ctx context.Context, name string) (*Movie, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) RegisterMovie(ctx context.Context, movie Movie) (int, error) {
	res, err := s.repo.FindMovieByNameAndLanguage(ctx, movie)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
	}
	if res != nil && err == nil {
		return 0, errors.New("this movie already exist")
	}
	movieId, err := s.repo.CreateMovie(ctx, movie)
	if err != nil {
		return 0, err
	}
	return movieId, nil
}

func (s *service) DeleteMovie(ctx context.Context, movieId int) error {
	err := s.repo.DeleteMovie(ctx, movieId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetMovieDetailsByID(ctx context.Context, movieId int) (*Movie, error) {
	movie, err := s.repo.GetMovieDetailsById(ctx, movieId)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *service) ListMovies(ctx context.Context) ([]Movie, error) {
	movies, err := s.repo.GetMovies(ctx)
	if err != nil {
		return nil, err
	}
	if len(movies) < 1 {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}

func (s *service) UpdateMovie(ctx context.Context, movie Movie, movieId int) error {
	err := s.repo.UpdateMovie(ctx, movie, movieId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetMovieByName(ctx context.Context, name string) (*Movie, error) {
	movie, err := s.repo.GetMovieByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *service) GetMoviesByGenre(ctx context.Context, genre string) ([]Movie, error) {
	movies, err := s.repo.GetMoviesByGenre(ctx, genre)
	if err != nil {
		return nil, err
	}
	if len(movies) < 1 {
		return nil, fmt.Errorf("no movies found in this genre %s", genre)
	}
	return movies, nil
}

func (s *service) GetMoviesByLanguage(ctx context.Context, language string) ([]Movie, error) {
	movies, err := s.repo.GetMoviesByLanguage(ctx, language)
	if err != nil {
		return nil, err
	}
	if len(movies) < 1 {
		return nil, fmt.Errorf("no movies found in this genre %s", language)
	}
	return movies, nil
}
