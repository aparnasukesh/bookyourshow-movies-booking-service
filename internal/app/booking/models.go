package booking

import (
	"time"
)

type Booking struct {
	BookingID     uint          `gorm:"primaryKey;autoIncrement" json:"booking_id"`
	UserID        uint          `gorm:"not null" json:"user_id"`
	ShowtimeID    uint          `gorm:"not null" json:"showtime_id"`
	BookingDate   time.Time     `gorm:"type:timestamp;not null" json:"booking_date"`
	TotalAmount   float64       `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	PaymentStatus string        `gorm:"type:varchar(50);not null" json:"payment_status"`
	BookingSeats  []BookingSeat `gorm:"foreignKey:BookingID" json:"booking_seats"`
}

type BookingSeat struct {
	BookingID uint `gorm:"primaryKey;autoIncrement:false" json:"booking_id"`
	SeatID    uint `gorm:"primaryKey;autoIncrement:false" json:"seat_id"`
}
