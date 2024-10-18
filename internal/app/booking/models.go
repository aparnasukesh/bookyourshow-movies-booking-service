package booking

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	BookingID     uint           `gorm:"primaryKey;autoIncrement" json:"booking_id"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	ShowtimeID    uint           `gorm:"not null" json:"showtime_id"`
	ScreenID      uint           `json:"screen_id"`
	BookingDate   time.Time      `gorm:"type:timestamp;not null" json:"booking_date"`
	TotalAmount   float64        `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	PaymentStatus string         `gorm:"type:varchar(50);not null" json:"payment_status"`
	BookingSeats  []BookingSeat  `gorm:"foreignKey:BookingID" json:"booking_seats"`
}

type BookingSeat struct {
	BookingID uint           `gorm:"primaryKey;autoIncrement:false" json:"booking_id"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	SeatID    uint           `gorm:"primaryKey;autoIncrement:false" json:"seat_id"`
}

type CreateBookingRequest struct {
	UserID      int     `json:"user_id"`
	ShowtimeID  int     `json:"showtime_id"`
	ScreenID    uint    `json:"screen_id"`
	SeatIDs     []int   `json:"seat_ids"`
	TotalAmount float64 `json:"total_amount"`
}
