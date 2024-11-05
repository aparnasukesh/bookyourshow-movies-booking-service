package booking

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateBooking(ctx context.Context, booking *Booking) error
	CreateBookingSeats(ctx context.Context, bookingSeats []BookingSeat) error
	GetBookingByID(ctx context.Context, bookingId int) (*Booking, error)
	ListBookingsByUser(ctx context.Context, userId int) ([]Booking, error)
	DeleteBookingByBookingID(ctx context.Context, bookingId int) error
	UpdateBookingStatusByBookingID(ctx context.Context, bookingId int, status string) error
	DeleteBookingSeats(ctx context.Context, bookingId int) error
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

func (r *repository) GetBookingByID(ctx context.Context, bookingId int) (*Booking, error) {
	booking := &Booking{}
	res := r.db.Preload("BookingSeats").Where("booking_id = ?", bookingId).First(&booking)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no booking found with id %d", bookingId)
		}
		return nil, res.Error
	}
	return booking, nil
}

func (r *repository) ListBookingsByUser(ctx context.Context, userId int) ([]Booking, error) {
	booking := []Booking{}
	res := r.db.Preload("BookingSeats").Where("user_id = ?", userId).Find(&booking)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no booking found with user id %d", userId)
		}
		return nil, res.Error
	}
	return booking, nil
}
func (r *repository) DeleteBookingByBookingID(ctx context.Context, bookingId int) error {
	if err := r.db.Where("booking_id = ?", bookingId).Delete(&Booking{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateBookingStatusByBookingID(ctx context.Context, bookingId int, status string) error {
	booking := Booking{}

	err := r.db.Model(&booking).Where("booking_id = ? AND payment_status ILIKE ?", bookingId, "Pending").Update("payment_status", status).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) DeleteBookingSeats(ctx context.Context, bookingId int) error {
	if err := r.db.Where("booking_id = ?", bookingId).Delete(&BookingSeat{}).Error; err != nil {
		return err
	}
	return nil
}
