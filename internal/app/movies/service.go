package movies

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type service struct {
	repo Repository
}

type Service interface {
	RegisterMovie(ctx context.Context, movie Movie) (int, error)
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
