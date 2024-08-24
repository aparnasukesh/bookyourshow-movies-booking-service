package theatres

import (
	"time"

	"github.com/aparnasukesh/movies-booking-svc/internal/app/movies"
	"gorm.io/gorm"
)

// Theatres
type Theater struct {
	gorm.Model
	Name            string          `json:"name"`
	Place           string          `json:"place"`
	City            string          `json:"city"`
	District        string          `json:"district"`
	State           string          `json:"state"`
	OwnerID         uint            `json:"owner_id"`
	NumberOfScreens int             `json:"number_of_screens"`
	TheaterTypeID   int             `json:"theater_type_id"`
	TheaterType     TheaterType     `gorm:"foreignKey:TheaterTypeID"`
	TheaterScreens  []TheaterScreen `gorm:"foreignKey:TheaterID"`
	MovieSchedules  []MovieSchedule `gorm:"foreignKey:TheaterID"`
}
type TheaterType struct {
	gorm.Model
	TheaterTypeName string    `json:"theater_type_name"`
	Theaters        []Theater `gorm:"foreignKey:TheaterTypeID"`
}

type ScreenType struct {
	gorm.Model
	ScreenTypeName string          `json:"screen_type_name"`
	TheaterScreens []TheaterScreen `gorm:"foreignKey:ScreenTypeID"`
}

type TheaterScreen struct {
	gorm.Model
	TheaterID    int        `json:"theater_id"`
	ScreenNumber int        `json:"screen_number"`
	SeatCapacity int        `json:"seat_capacity"`
	ScreenTypeID int        `json:"screen_type_id"`
	Theater      Theater    `gorm:"foreignKey:TheaterID"`
	ScreenType   ScreenType `gorm:"foreignKey:ScreenTypeID"`
	Seats        []Seat     `gorm:"foreignKey:ScreenID"`
	Showtimes    []Showtime `gorm:"foreignKey:ScreenID"`
}

type Showtime struct {
	gorm.Model
	MovieID       int           `json:"movie_id"`
	ScreenID      int           `json:"screen_id"`
	ShowDate      time.Time     `json:"show_date"`
	ShowTime      time.Time     `json:"show_time"`
	Movie         movies.Movie  `gorm:"foreignKey:MovieID"`
	TheaterScreen TheaterScreen `gorm:"foreignKey:ScreenID"`
}

type MovieSchedule struct {
	gorm.Model
	MovieID    int          `json:"movie_id"`
	TheaterID  int          `json:"theater_id"`
	ShowtimeID int          `json:"showtime_id"`
	Movie      movies.Movie `gorm:"foreignKey:MovieID"`
	Theater    Theater      `gorm:"foreignKey:TheaterID"`
	Showtime   Showtime     `gorm:"foreignKey:ShowtimeID"`
}

type SeatCategory struct {
	gorm.Model
	SeatCategoryName string `json:"seat_category_name"`
	Seats            []Seat `gorm:"foreignKey:SeatCategoryID"`
}

type Seat struct {
	gorm.Model
	ScreenID          int           `json:"screen_id"`
	SeatNumber        string        `json:"seat_number"`
	Row               string        `json:"row"`
	Column            int           `json:"column"`
	SeatCategoryID    int           `json:"seat_category_id"`
	SeatCategoryPrice float64       `json:"seat_category_price"`
	TheaterScreen     TheaterScreen `gorm:"foreignKey:ScreenID"`
	SeatCategory      SeatCategory  `gorm:"foreignKey:SeatCategoryID"`
}
