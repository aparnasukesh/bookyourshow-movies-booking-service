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
	// theater type
	AddTheaterType(ctx context.Context, theaterType TheaterType) error
	DeleteTheaterTypeByID(ctx context.Context, id int) error
	DeleteTheaterTypeByName(ctx context.Context, name string) error
	GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error)
	GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error)
	UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error
	ListTheaterTypes(ctx context.Context) ([]TheaterType, error)
	//screen types
	AddScreenType(ctx context.Context, screenType ScreenType) error
	DeleteScreenTypeByID(ctx context.Context, id int) error
	DeleteScreenTypeByName(ctx context.Context, name string) error
	GetScreenTypeByID(ctx context.Context, id int) (*ScreenType, error)
	GetScreenTypeByName(ctx context.Context, name string) (*ScreenType, error)
	UpdateScreenType(ctx context.Context, id int, screenType ScreenType) error
	ListScreenTypes(ctx context.Context) ([]ScreenType, error)
	// Seat category
	AddSeatCategory(ctx context.Context, seatCategory SeatCategory) error
	DeleteSeatCategoryByID(ctx context.Context, id int) error
	DeleteSeatCategoryByName(ctx context.Context, name string) error
	GetSeatCategoryByID(ctx context.Context, id int) (*SeatCategory, error)
	GetSeatCategoryByName(ctx context.Context, name string) (*SeatCategory, error)
	UpdateSeatCategory(ctx context.Context, id int, seatCategory SeatCategory) error
	ListSeatCategories(ctx context.Context) ([]SeatCategory, error)
	// Theater
	AddTheater(ctx context.Context, theater Theater) error
	DeleteTheaterByID(ctx context.Context, id int) error
	DeleteTheaterByName(ctx context.Context, name string) error
	GetTheaterByID(ctx context.Context, id int) (*Theater, error)
	GetTheaterByName(ctx context.Context, name string) (*Theater, error)
	UpdateTheater(ctx context.Context, id int, theater Theater) error
	ListTheaters(ctx context.Context) ([]Theater, error)
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

// screen type
func (s *service) AddScreenType(ctx context.Context, screenType ScreenType) error {
	res, err := s.repo.FindScreenTypeByName(ctx, screenType.ScreenTypeName)
	if res != nil && err == nil {
		return errors.New("screen type already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if err := s.repo.CreateScreenType(ctx, screenType); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteScreenTypeByID(ctx context.Context, id int) error {
	if err := s.repo.DeleteScreenTypeByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteScreenTypeByName(ctx context.Context, name string) error {
	if err := s.repo.DeleteScreenTypeByName(ctx, name); err != nil {
		return err
	}
	return nil
}

func (s *service) GetScreenTypeByID(ctx context.Context, id int) (*ScreenType, error) {
	screenType, err := s.repo.GetScreenTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return screenType, nil
}

func (s *service) GetScreenTypeByName(ctx context.Context, name string) (*ScreenType, error) {
	screenType, err := s.repo.GetScreenTypeByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return screenType, nil
}

func (s *service) ListScreenTypes(ctx context.Context) ([]ScreenType, error) {
	screenTypes, err := s.repo.ListScreenTypes(ctx)
	if err != nil {
		return nil, err
	}
	return screenTypes, nil
}

func (s *service) UpdateScreenType(ctx context.Context, id int, screenType ScreenType) error {
	err := s.repo.UpdateScreenType(ctx, id, screenType)
	if err != nil {
		return err
	}
	return nil
}

// seat category
func (s *service) AddSeatCategory(ctx context.Context, seatCategory SeatCategory) error {
	res, err := s.repo.FindSeatCategoryByName(ctx, seatCategory.SeatCategoryName)
	if res != nil && err == nil {
		return errors.New("seat category already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if err := s.repo.CreateSeatCategory(ctx, seatCategory); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteSeatCategoryByID(ctx context.Context, id int) error {
	if err := s.repo.DeleteSeatCategoryByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteSeatCategoryByName(ctx context.Context, name string) error {
	if err := s.repo.DeleteSeatCategoryByName(ctx, name); err != nil {
		return err
	}
	return nil
}

func (s *service) GetSeatCategoryByID(ctx context.Context, id int) (*SeatCategory, error) {
	seatCategory, err := s.repo.GetSeatCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return seatCategory, nil
}

func (s *service) GetSeatCategoryByName(ctx context.Context, name string) (*SeatCategory, error) {
	seatCategory, err := s.repo.GetSeatCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return seatCategory, nil
}

func (s *service) UpdateSeatCategory(ctx context.Context, id int, seatCategory SeatCategory) error {
	err := s.repo.UpdateSeatCategory(ctx, id, seatCategory)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListSeatCategories(ctx context.Context) ([]SeatCategory, error) {
	seatCategories, err := s.repo.ListSeatCategories(ctx)
	if err != nil {
		return nil, err
	}
	return seatCategories, nil
}

// Theater
func (s *service) AddTheater(ctx context.Context, theater Theater) error {
	res, err := s.repo.FindTheaterByNameAndOwnerId(ctx, theater.Name, theater.OwnerID)
	if res != nil && err == nil {
		return errors.New("theater already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if err := s.repo.CreateTheater(ctx, theater); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTheaterByID(ctx context.Context, id int) error {
	if err := s.repo.DeleteTheaterByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTheaterByName(ctx context.Context, name string) error {
	if err := s.repo.DeleteTheaterByName(ctx, name); err != nil {
		return err
	}
	return nil
}

func (s *service) GetTheaterByID(ctx context.Context, id int) (*Theater, error) {
	theater, err := s.repo.GetTheaterByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return theater, nil
}

func (s *service) GetTheaterByName(ctx context.Context, name string) (*Theater, error) {
	theater, err := s.repo.GetTheaterByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return theater, nil
}

func (s *service) UpdateTheater(ctx context.Context, id int, theater Theater) error {
	err := s.repo.UpdateTheater(ctx, id, theater)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListTheaters(ctx context.Context) ([]Theater, error) {
	theaters, err := s.repo.ListTheaters(ctx)
	if err != nil {
		return nil, err
	}
	return theaters, nil
}
