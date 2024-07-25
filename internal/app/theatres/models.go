package theatres

import (
	"time"

	"gorm.io/gorm"
)

type Theater struct {
	gorm.Model
	Name            string          `gorm:"type:varchar(100);not null"`
	Location        string          `gorm:"type:varchar(255);not null"`
	NumberOfScreens int             `gorm:"not null"`
	OwnerID         uint            `gorm:"not null"`
	Screens         []Screen        `gorm:"foreignKey:TheaterID"`
	MovieSchedules  []MovieSchedule `gorm:"foreignKey:TheaterID"`
}

type Screen struct {
	gorm.Model
	TheaterID    uint       `gorm:"not null"`
	ScreenNumber int        `gorm:"not null"`
	SeatCapacity int        `gorm:"not null"`
	ScreenTypeID uint       `gorm:"not null"`
	Showtimes    []Showtime `gorm:"foreignKey:ScreenID"`
	Seats        []Seat     `gorm:"foreignKey:ScreenID"`
}

type Showtime struct {
	gorm.Model
	MovieID        uint            `gorm:"not null"`
	ScreenID       uint            `gorm:"not null"`
	ShowDate       time.Time       `gorm:"not null"`
	ShowTime       time.Time       `gorm:"not null"`
	MovieSchedules []MovieSchedule `gorm:"foreignKey:ShowtimeID"`
}

type MovieSchedule struct {
	gorm.Model
	MovieID    uint `gorm:"not null"`
	TheaterID  uint `gorm:"not null"`
	ShowtimeID uint `gorm:"not null"`
}

type Seat struct {
	gorm.Model
	ScreenID       uint   `gorm:"not null"`
	SeatNumber     int    `gorm:"not null"`
	Row            string `gorm:"type:varchar(10);not null"`
	Column         string `gorm:"type:varchar(10);not null"`
	SeatCategoryID uint   `gorm:"not null"`
}

type SeatCategory struct {
	gorm.Model
	SeatCategoryName  string  `gorm:"type:varchar(50);not null"`
	SeatCategoryPrice float64 `gorm:"type:decimal(10,2);not null"`
	Seats             []Seat  `gorm:"foreignKey:SeatCategoryID"`
}
