package theatres

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type service struct {
	repo Repository
}

type Service interface {
	AddTheaterType(ctx context.Context, theaterType TheaterType) error
	DeleteTheaterTypeByID(ctx context.Context, id int) error
	DeleteTheaterTypeByName(ctx context.Context, name string) error
	GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error)
	GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error)
	UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error
	ListTheaterTypes(ctx context.Context) ([]TheaterType, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Theater-Type
func (s *service) AddTheaterType(ctx context.Context, theaterType TheaterType) error {
	res, err := s.repo.FindTheatertypeByName(ctx, theaterType.TheaterTypeName)
	if res != nil && err == nil {
		return errors.New("theater type already exist")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if err := s.repo.CreateTheaterType(ctx, theaterType); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTheaterTypeByID(ctx context.Context, id int) error {
	if err := s.repo.DeleteTheaterTypeByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTheaterTypeByName(ctx context.Context, name string) error {
	if err := s.repo.DeleteTheaterTypeByName(ctx, name); err != nil {
		return err
	}
	return nil
}

func (s *service) GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error) {
	theaterType, err := s.repo.GetTheaterTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return theaterType, nil
}

func (s *service) GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error) {
	theaterType, err := s.repo.GetTheaterTypeByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return theaterType, nil
}

func (s *service) ListTheaterTypes(ctx context.Context) ([]TheaterType, error) {
	theaterTypes, err := s.repo.ListTheaterTypes(ctx)
	if err != nil {
		return nil, err
	}
	return theaterTypes, nil
}

func (s *service) UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error {
	err := s.repo.UpdateTheaterType(ctx, id, theaterType)
	if err != nil {
		return err
	}
	return nil
}
