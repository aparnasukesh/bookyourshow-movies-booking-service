package booking

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateBooking(ctx context.Context, booking *Booking) error
	CreateBookingSeats(ctx context.Context, bookingSeats []BookingSeat) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateBooking(ctx context.Context, booking *Booking) error {
	if err := r.db.Create(&booking).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateBookingSeats(ctx context.Context, bookingSeats []BookingSeat) error {
	if err := r.db.Create(&bookingSeats).Error; err != nil {
		return err
	}
	return nil
}
