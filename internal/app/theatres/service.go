package theatres

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	GetTheaterByName(ctx context.Context, name string) ([]Theater, error)
	UpdateTheater(ctx context.Context, id int, theater Theater) error
	ListTheaters(ctx context.Context) ([]Theater, error)
	//Theater screen
	AddTheaterScreen(ctx context.Context, theaterScreen TheaterScreen) error
	DeleteTheaterScreenByID(ctx context.Context, id int) error
	DeleteTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) error
	GetTheaterScreenByID(ctx context.Context, id int) (*TheaterScreen, error)
	GetTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) (*TheaterScreen, error)
	UpdateTheaterScreen(ctx context.Context, id int, theaterScreen TheaterScreen) error
	ListTheaterScreens(ctx context.Context, theaterId int) ([]TheaterScreen, error)
	//Show time
	AddShowtime(ctx context.Context, showtime Showtime) error
	DeleteShowtimeByID(ctx context.Context, id int) error
	DeleteShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) error
	GetShowtimeByID(ctx context.Context, id int) (*Showtime, error)
	GetShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) (*Showtime, error)
	UpdateShowtime(ctx context.Context, id int, showtime Showtime) error
	ListShowtimes(ctx context.Context, movieID int) ([]Showtime, error)
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
	if len(theaterTypes) < 1 {
		return nil, errors.New("theater types are not found")
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
	if len(screenTypes) < 1 {
		return nil, errors.New("no screen types found")
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
	if len(seatCategories) < 1 {
		return nil, errors.New("no seat categories found")
	}
	return seatCategories, nil
}

// Theater
func (s *service) AddTheater(ctx context.Context, theater Theater) error {
	stateCount, err := s.repo.CountTheatersByOwnerAndState(ctx, theater.OwnerID, theater.State)
	if err != nil {
		return fmt.Errorf("failed to count theaters for owner in the state: %w", err)
	}
	if MaxTheatersPerOwnerInState <= stateCount {
		return errors.New("the owner has reached the maximum limit of theaters in this state")
	}
	districtCount, err := s.repo.CountTheatersByOwnerAndDistrict(ctx, theater.OwnerID, theater.State)
	if err != nil {
		return fmt.Errorf("failed to count theaters for owner in the state: %w", err)
	}
	if MaxTheatersPerOwnerInDistrict <= districtCount {
		return errors.New("the owner has reached the maximum limit of theaters in this district")
	}
	cityCount, err := s.repo.CountTheatersByOwnerAndCity(ctx, theater.OwnerID, theater.City)
	if err != nil {
		return fmt.Errorf("failed to count theaters for owner in the city: %w", err)
	}
	if MaxTheatersPerOwnerInCity <= cityCount {
		return errors.New("the owner has reached the maximum limit of theaters in this city")
	}
	placeCount, err := s.repo.CountTheatersByOwnerAndPlace(ctx, theater.OwnerID, theater.Place)
	if err != nil {
		return fmt.Errorf("failed to count theaters for owner in the place: %w", err)
	}
	if MaxTheatersPerOwnerInPlace <= placeCount {
		return errors.New("the owner has reached the maximum limit of theaters in this place")
	}
	res, err := s.repo.FindTheaterByNamePlaceAndOwnerId(ctx, theater.Name, theater.Place, theater.OwnerID)
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

func (s *service) GetTheaterByName(ctx context.Context, name string) ([]Theater, error) {
	theaters, err := s.repo.GetTheaterByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if len(theaters) < 1 {
		return nil, fmt.Errorf("no theaters with name %s", name)
	}
	return theaters, nil
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
	if len(theaters) < 1 {
		return nil, errors.New("no theaters found")
	}
	return theaters, nil
}

// Theater Screens
func (s *service) AddTheaterScreen(ctx context.Context, theaterScreen TheaterScreen) error {
	theater, err := s.repo.GetTheaterByID(ctx, theaterScreen.TheaterID)
	if theater == nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("theater not exist with theater id %d", theaterScreen.TheaterID)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	screen, err := s.repo.GetScreenTypeByID(ctx, theaterScreen.ScreenTypeID)
	if screen == nil && err == gorm.ErrRecordNotFound {
		return fmt.Errorf("sceen type not exist with screen type id %d", theaterScreen.ScreenTypeID)
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	res, err := s.repo.FindTheaterScreenByTheaterIDAndScreenNumber(ctx, theaterScreen.TheaterID, theaterScreen.ScreenNumber)
	if res != nil && err == nil {
		return errors.New("theater screen already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if MaxScreenPerTheater < theaterScreen.ScreenNumber {
		return errors.New("the theater has reached the maximum screen limit")
	}
	if err := s.repo.CreateTheaterScreen(ctx, theaterScreen); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTheaterScreenByID(ctx context.Context, id int) error {
	if err := s.repo.DeleteTheaterScreenByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) error {
	if err := s.repo.DeleteTheaterScreenByNumber(ctx, theaterID, screenNumber); err != nil {
		return err
	}
	return nil
}

func (s *service) GetTheaterScreenByID(ctx context.Context, id int) (*TheaterScreen, error) {
	theaterScreen, err := s.repo.GetTheaterScreenByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return theaterScreen, nil
}

func (s *service) GetTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) (*TheaterScreen, error) {
	theaterScreen, err := s.repo.GetTheaterScreenByNumber(ctx, theaterID, screenNumber)
	if err != nil {
		return nil, err
	}
	return theaterScreen, nil
}

func (s *service) UpdateTheaterScreen(ctx context.Context, id int, theaterScreen TheaterScreen) error {
	err := s.repo.UpdateTheaterScreen(ctx, id, theaterScreen)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListTheaterScreens(ctx context.Context, theaterId int) ([]TheaterScreen, error) {
	theaterScreens, err := s.repo.ListTheaterScreens(ctx, theaterId)
	if err != nil {
		return nil, err
	}
	if len(theaterScreens) < 1 {
		return nil, errors.New("no theater screens found")
	}
	return theaterScreens, nil
}

// Showtimes
func (s *service) AddShowtime(ctx context.Context, showtime Showtime) error {
	res, err := s.repo.FindShowtimeByDetails(ctx, showtime.MovieID, showtime.ScreenID, showtime.ShowDate, showtime.ShowTime)
	if res != nil && err == nil {
		return errors.New("showtime already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}

	if err := s.repo.CreateShowtime(ctx, showtime); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteShowtimeByID(ctx context.Context, id int) error {
	if err := s.repo.DeleteShowtimeByID(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) error {
	if err := s.repo.DeleteShowtimeByDetails(ctx, movieID, screenID, showDate, showTime); err != nil {
		return err
	}
	return nil
}

func (s *service) GetShowtimeByID(ctx context.Context, id int) (*Showtime, error) {
	showtime, err := s.repo.GetShowtimeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return showtime, nil
}

func (s *service) GetShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) (*Showtime, error) {
	showtime, err := s.repo.GetShowtimeByDetails(ctx, movieID, screenID, showDate, showTime)
	if err != nil {
		return nil, err
	}
	return showtime, nil
}

func (s *service) UpdateShowtime(ctx context.Context, id int, showtime Showtime) error {
	err := s.repo.UpdateShowtime(ctx, id, showtime)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListShowtimes(ctx context.Context, movieID int) ([]Showtime, error) {
	showtimes, err := s.repo.ListShowtimes(ctx, movieID)
	if err != nil {
		return nil, err
	}
	if len(showtimes) < 1 {
		return nil, errors.New("no showtimes found")
	}
	return showtimes, nil
}
