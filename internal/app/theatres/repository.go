package theatres

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	//theater type
	CreateTheaterType(ctx context.Context, theaterType TheaterType) error
	DeleteTheaterTypeByID(ctx context.Context, id int) error
	DeleteTheaterTypeByName(ctx context.Context, name string) error
	FindTheatertypeByName(ctx context.Context, name string) (*TheaterType, error)
	GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error)
	GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error)
	UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error
	ListTheaterTypes(ctx context.Context) ([]TheaterType, error)
	//screen type
	CreateScreenType(ctx context.Context, screenType ScreenType) error
	DeleteScreenTypeByID(ctx context.Context, id int) error
	DeleteScreenTypeByName(ctx context.Context, name string) error
	FindScreenTypeByName(ctx context.Context, name string) (*ScreenType, error)
	GetScreenTypeByID(ctx context.Context, id int) (*ScreenType, error)
	GetScreenTypeByName(ctx context.Context, name string) (*ScreenType, error)
	UpdateScreenType(ctx context.Context, id int, screenType ScreenType) error
	ListScreenTypes(ctx context.Context) ([]ScreenType, error)
	//seat category
	CreateSeatCategory(ctx context.Context, seatCategory SeatCategory) error
	DeleteSeatCategoryByID(ctx context.Context, id int) error
	DeleteSeatCategoryByName(ctx context.Context, name string) error
	FindSeatCategoryByName(ctx context.Context, name string) (*SeatCategory, error)
	GetSeatCategoryByID(ctx context.Context, id int) (*SeatCategory, error)
	GetSeatCategoryByName(ctx context.Context, name string) (*SeatCategory, error)
	UpdateSeatCategory(ctx context.Context, id int, seatCategory SeatCategory) error
	ListSeatCategories(ctx context.Context) ([]SeatCategory, error)
	//Theater
	CreateTheater(ctx context.Context, theater Theater) error
	DeleteTheaterByID(ctx context.Context, id int) error
	DeleteTheaterByName(ctx context.Context, name string) error
	FindTheaterByNameAndOwnerId(ctx context.Context, theaterName string, theaterOwnerID uint) (*Theater, error)
	GetTheaterByID(ctx context.Context, id int) (*Theater, error)
	GetTheaterByName(ctx context.Context, name string) (*Theater, error)
	UpdateTheater(ctx context.Context, id int, theater Theater) error
	ListTheaters(ctx context.Context) ([]Theater, error)
	//Theater screen
	CreateTheaterScreen(ctx context.Context, theaterScreen TheaterScreen) error
	DeleteTheaterScreenByID(ctx context.Context, id int) error
	DeleteTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) error
	FindTheaterScreenByTheaterIDAndScreenNumber(ctx context.Context, theaterID int, screenNumber int) (*TheaterScreen, error)
	GetTheaterScreenByID(ctx context.Context, id int) (*TheaterScreen, error)
	GetTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) (*TheaterScreen, error)
	UpdateTheaterScreen(ctx context.Context, id int, theaterScreen TheaterScreen) error
	ListTheaterScreens(ctx context.Context, theaterId int) ([]TheaterScreen, error)
	//Show Time
	FindShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) (*Showtime, error)
	CreateShowtime(ctx context.Context, showtime Showtime) error
	DeleteShowtimeByID(ctx context.Context, id int) error
	DeleteShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) error
	GetShowtimeByID(ctx context.Context, id int) (*Showtime, error)
	GetShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) (*Showtime, error)
	ListShowtimes(ctx context.Context, movieID int) ([]Showtime, error)
	UpdateShowtime(ctx context.Context, id int, showtime Showtime) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindTheatertypeByName(ctx context.Context, name string) (*TheaterType, error) {
	theaterType := &TheaterType{}
	res := r.db.Where("theater_type_name ILIKE ?", name).First(&theaterType)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		} else if res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return theaterType, nil

}

func (r *repository) CreateTheaterType(ctx context.Context, theaterType TheaterType) error {
	if err := r.db.Create(&theaterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterTypeByID(ctx context.Context, id int) error {
	theaterType := &TheaterType{}
	if err := r.db.Where("id =?", id).Delete(&theaterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterTypeByName(ctx context.Context, name string) error {
	theaterType := &TheaterType{}
	if err := r.db.Where("theater_type_name ILIKE ?", name).Delete(&theaterType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTheaterTypeByID(ctx context.Context, id int) (*TheaterType, error) {
	theatertype := TheaterType{}
	if err := r.db.Where("id =?", id).First(&theatertype).Error; err != nil {
		return nil, err
	}
	return &theatertype, nil
}

func (r *repository) GetTheaterTypeByName(ctx context.Context, name string) (*TheaterType, error) {
	theaterType := &TheaterType{}
	if err := r.db.Where("theater_type_name ILIKE ?", name).First(&theaterType).Error; err != nil {
		return nil, err
	}
	return theaterType, nil
}

func (r *repository) ListTheaterTypes(ctx context.Context) ([]TheaterType, error) {
	theaterTypes := []TheaterType{}
	if err := r.db.Find(&theaterTypes).Error; err != nil {
		return nil, err
	}
	return theaterTypes, nil
}

func (r *repository) UpdateTheaterType(ctx context.Context, id int, theaterType TheaterType) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&TheaterType{}).Where("id = ?", id).Updates(theaterType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// screen types
func (r *repository) FindScreenTypeByName(ctx context.Context, name string) (*ScreenType, error) {
	screenType := &ScreenType{}
	res := r.db.Where("screen_type_name ILIKE ?", name).First(&screenType)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		} else if res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return screenType, nil
}

func (r *repository) CreateScreenType(ctx context.Context, screenType ScreenType) error {
	if err := r.db.Create(&screenType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteScreenTypeByID(ctx context.Context, id int) error {
	screenType := &ScreenType{}
	if err := r.db.Where("id = ?", id).Delete(&screenType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteScreenTypeByName(ctx context.Context, name string) error {
	screenType := &ScreenType{}
	if err := r.db.Where("screen_type_name ILIKE ?", name).Delete(&screenType).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetScreenTypeByID(ctx context.Context, id int) (*ScreenType, error) {
	screenType := ScreenType{}
	if err := r.db.Where("id = ?", id).First(&screenType).Error; err != nil {
		return nil, err
	}
	return &screenType, nil
}

func (r *repository) GetScreenTypeByName(ctx context.Context, name string) (*ScreenType, error) {
	screenType := &ScreenType{}
	if err := r.db.Where("screen_type_name ILIKE ?", name).First(&screenType).Error; err != nil {
		return nil, err
	}
	return screenType, nil
}

func (r *repository) ListScreenTypes(ctx context.Context) ([]ScreenType, error) {
	screenTypes := []ScreenType{}
	if err := r.db.Find(&screenTypes).Error; err != nil {
		return nil, err
	}
	return screenTypes, nil
}

func (r *repository) UpdateScreenType(ctx context.Context, id int, screenType ScreenType) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&ScreenType{}).Where("id = ?", id).Updates(screenType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// seat category
func (r *repository) FindSeatCategoryByName(ctx context.Context, name string) (*SeatCategory, error) {
	seatCategory := &SeatCategory{}
	res := r.db.Where("seat_category_name ILIKE ?", name).First(&seatCategory)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		} else if res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return seatCategory, nil
}

func (r *repository) CreateSeatCategory(ctx context.Context, seatCategory SeatCategory) error {
	if err := r.db.Create(&seatCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteSeatCategoryByID(ctx context.Context, id int) error {
	seatCategory := &SeatCategory{}
	if err := r.db.Where("id = ?", id).Delete(&seatCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteSeatCategoryByName(ctx context.Context, name string) error {
	seatCategory := &SeatCategory{}
	if err := r.db.Where("seat_category_name ILIKE ?", name).Delete(&seatCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetSeatCategoryByID(ctx context.Context, id int) (*SeatCategory, error) {
	seatCategory := SeatCategory{}
	if err := r.db.Where("id = ?", id).First(&seatCategory).Error; err != nil {
		return nil, err
	}
	return &seatCategory, nil
}

func (r *repository) GetSeatCategoryByName(ctx context.Context, name string) (*SeatCategory, error) {
	seatCategory := &SeatCategory{}
	if err := r.db.Where("seat_category_name ILIKE ?", name).First(&seatCategory).Error; err != nil {
		return nil, err
	}
	return seatCategory, nil
}

func (r *repository) ListSeatCategories(ctx context.Context) ([]SeatCategory, error) {
	seatCategories := []SeatCategory{}
	if err := r.db.Find(&seatCategories).Error; err != nil {
		return nil, err
	}
	return seatCategories, nil
}

func (r *repository) UpdateSeatCategory(ctx context.Context, id int, seatCategory SeatCategory) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&SeatCategory{}).Where("id = ?", id).Updates(seatCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Theater
func (r *repository) FindTheaterByNameAndOwnerId(ctx context.Context, theaterName string, theaterOwnerID uint) (*Theater, error) {
	theater := &Theater{}
	res := r.db.Where("name ILIKE ? AND owner_id = ?", theaterName, theaterOwnerID).First(theater)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound || res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return theater, nil
}

func (r *repository) CreateTheater(ctx context.Context, theater Theater) error {
	if err := r.db.Create(&theater).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterByID(ctx context.Context, id int) error {
	theater := &Theater{}
	if err := r.db.Where("id = ?", id).Delete(&theater).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterByName(ctx context.Context, name string) error {
	theater := &Theater{}
	if err := r.db.Where("name ILIKE ?", name).Delete(&theater).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTheaterByID(ctx context.Context, id int) (*Theater, error) {
	theater := Theater{}
	if err := r.db.Where("id = ?", id).First(&theater).Error; err != nil {
		return nil, err
	}
	return &theater, nil
}

func (r *repository) GetTheaterByName(ctx context.Context, name string) (*Theater, error) {
	theater := &Theater{}
	if err := r.db.Where("name ILIKE ?", name).First(&theater).Error; err != nil {
		return nil, err
	}
	return theater, nil
}

func (r *repository) ListTheaters(ctx context.Context) ([]Theater, error) {
	theaters := []Theater{}
	if err := r.db.Find(&theaters).Error; err != nil {
		return nil, err
	}
	return theaters, nil
}

func (r *repository) UpdateTheater(ctx context.Context, id int, theater Theater) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&Theater{}).Where("id = ?", id).Updates(theater)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// TheaterScreen
func (r *repository) FindTheaterScreenByTheaterIDAndScreenNumber(ctx context.Context, theaterID int, screenNumber int) (*TheaterScreen, error) {
	theaterScreen := &TheaterScreen{}
	res := r.db.Where("theater_id = ? AND screen_number = ?", theaterID, screenNumber).First(theaterScreen)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound || res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return theaterScreen, nil
}

func (r *repository) CreateTheaterScreen(ctx context.Context, theaterScreen TheaterScreen) error {
	if err := r.db.Create(&theaterScreen).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterScreenByID(ctx context.Context, id int) error {
	theaterScreen := &TheaterScreen{}
	if err := r.db.Where("id = ?", id).Delete(&theaterScreen).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) error {
	theaterScreen := &TheaterScreen{}
	if err := r.db.Where("theater_id = ? AND screen_number = ?", theaterID, screenNumber).Delete(&theaterScreen).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTheaterScreenByID(ctx context.Context, id int) (*TheaterScreen, error) {
	theaterScreen := &TheaterScreen{}
	if err := r.db.Where("id = ?", id).First(&theaterScreen).Error; err != nil {
		return nil, err
	}
	return theaterScreen, nil
}

func (r *repository) GetTheaterScreenByNumber(ctx context.Context, theaterID int, screenNumber int) (*TheaterScreen, error) {
	theaterScreen := &TheaterScreen{}
	if err := r.db.Where("theater_id = ? AND screen_number = ?", theaterID, screenNumber).First(&theaterScreen).Error; err != nil {
		return nil, err
	}
	return theaterScreen, nil
}

func (r *repository) ListTheaterScreens(ctx context.Context, theaterId int) ([]TheaterScreen, error) {
	theaterScreens := []TheaterScreen{}
	if err := r.db.Where("theater_id =?", theaterId).Find(&theaterScreens).Error; err != nil {
		return nil, err
	}
	return theaterScreens, nil
}

func (r *repository) UpdateTheaterScreen(ctx context.Context, id int, theaterScreen TheaterScreen) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&TheaterScreen{}).Where("id = ?", id).Updates(theaterScreen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Show time
func (r *repository) FindShowtimeByMovieIDAndScreenID(ctx context.Context, movieID int, screenID int) (*Showtime, error) {
	showtime := &Showtime{}
	res := r.db.Where("movie_id = ? AND screen_id = ?", movieID, screenID).First(showtime)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound || res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return showtime, nil
}

func (r *repository) CreateShowtime(ctx context.Context, showtime Showtime) error {
	if err := r.db.Create(&showtime).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteShowtimeByID(ctx context.Context, id int) error {
	showtime := &Showtime{}
	if err := r.db.Where("id = ?", id).Delete(&showtime).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) error {
	showtime := &Showtime{}
	if err := r.db.Where("movie_id = ? AND screen_id = ? AND show_date = ? AND show_time = ?", movieID, screenID, showDate, showTime).Delete(&showtime).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetShowtimeByID(ctx context.Context, id int) (*Showtime, error) {
	showtime := &Showtime{}
	if err := r.db.Where("id = ?", id).First(&showtime).Error; err != nil {
		return nil, err
	}
	return showtime, nil
}

func (r *repository) GetShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) (*Showtime, error) {
	showtime := &Showtime{}
	if err := r.db.Where("movie_id = ? AND screen_id = ? AND show_date = ? AND show_time = ?", movieID, screenID, showDate, showTime).First(&showtime).Error; err != nil {
		return nil, err
	}
	return showtime, nil
}

func (r *repository) ListShowtimes(ctx context.Context, movieID int) ([]Showtime, error) {
	showtimes := []Showtime{}
	if err := r.db.Where("movie_id = ?", movieID).Find(&showtimes).Error; err != nil {
		return nil, err
	}
	return showtimes, nil
}

func (r *repository) UpdateShowtime(ctx context.Context, id int, showtime Showtime) error {
	r.db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	result := r.db.Model(&Showtime{}).Where("id = ?", id).Updates(showtime)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) FindShowtimeByDetails(ctx context.Context, movieID int, screenID int, showDate time.Time, showTime time.Time) (*Showtime, error) {
	showtime := &Showtime{}
	res := r.db.Where("movie_id = ? AND screen_id = ? AND show_date = ? AND show_time = ?", movieID, screenID, showDate, showTime).First(showtime)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound || res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}
	return showtime, nil
}
