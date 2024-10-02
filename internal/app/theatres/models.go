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

type TheaterUpdateInput struct {
	Name            string `json:"name,omitempty"`
	Place           string `json:"place,omitempty"`
	City            string `json:"city,omitempty"`
	District        string `json:"district,omitempty"`
	State           string `json:"state,omitempty"`
	OwnerID         uint   `json:"owner_id,omitempty"`
	NumberOfScreens int    `json:"number_of_screens,omitempty"`
	TheaterTypeID   int    `json:"theater_type_id,omitempty"`
}

// Theater Type
type TheaterType struct {
	gorm.Model
	TheaterTypeName string    `json:"theater_type_name"`
	Theaters        []Theater `gorm:"foreignKey:TheaterTypeID"`
}

// Screen Type
type ScreenType struct {
	gorm.Model
	ScreenTypeName string          `json:"screen_type_name"`
	TheaterScreens []TheaterScreen `gorm:"foreignKey:ScreenTypeID"`
}

// Theater Screen
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

// Showtime
type Showtime struct {
	gorm.Model
	MovieID       int           `json:"movie_id"`
	ScreenID      int           `json:"screen_id"`
	ShowDate      time.Time     `json:"show_date"`
	ShowTime      time.Time     `json:"show_time"`
	Movie         Movie         `gorm:"foreignKey:MovieID"`
	TheaterScreen TheaterScreen `gorm:"foreignKey:ScreenID"`
}

// Movie Schedule
type MovieSchedule struct {
	gorm.Model
	MovieID    int          `json:"movie_id"`
	TheaterID  int          `json:"theater_id"`
	ShowtimeID int          `json:"showtime_id"`
	Movie      movies.Movie `gorm:"foreignKey:MovieID"`
	Theater    Theater      `gorm:"foreignKey:TheaterID"`
	Showtime   Showtime     `gorm:"foreignKey:ShowtimeID"`
}

// Seat Category
type SeatCategory struct {
	gorm.Model
	SeatCategoryName string `json:"seat_category_name"`
	Seats            []Seat `gorm:"foreignKey:SeatCategoryID"`
}

// Seat
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

type RowSeatCategoryPrice struct {
	RowStart          string  `json:"row_start"`
	RowEnd            string  `json:"row_end"`
	SeatCategoryId    int     `json:"seat_category_id"`
	SeatCategoryPrice float32 `json:"seat_category_price"`
}

type CreateSeatsRequest struct {
	ID           int                    `json:"id"`
	ScreenId     int                    `json:"screen_id"`
	TotalRows    int                    `json:"total_rows"`
	TotalColumns int                    `json:"total_columns"`
	SeatRequest  []RowSeatCategoryPrice `json:"seat_request"`
}
type TheaterWithTypeResponse struct {
	ID              int                 `json:"id"`
	Name            string              `json:"name"`
	Place           string              `json:"place"`
	City            string              `json:"city"`
	District        string              `json:"district"`
	State           string              `json:"state"`
	OwnerID         int                 `json:"owner_id"`
	NumberOfScreens int                 `json:"number_of_screens"`
	TheaterType     TheaterTypeResponse `json:"TheaterType"`
}

type TheaterTypeResponse struct {
	ID              int    `json:"id"`
	TheaterTypeName string `json:"theater_type_name"`
}

type Movie struct {
	ID          uint      `json:"id"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Duration    int       `gorm:"not null"`
	Genre       string    `gorm:"type:varchar(50)"`
	ReleaseDate time.Time `gorm:"not null"`
	Rating      float64   `gorm:"type:decimal(3,1)"`
	Language    string    `gorm:"type:varchar(100);not null"`
}
