package booking

import (
	"context"
	"fmt"
	"time"

	"github.com/aparnasukesh/movies-booking-svc/internal/app/movies"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/theatres"
	"gorm.io/gorm"
)

type service struct {
	db          *gorm.DB
	repo        Repository
	movieRepo   movies.Repository
	theaterRepo theatres.Repository
}

type Service interface {
	CreateBooking(ctx context.Context, createReq CreateBookingRequest) (*Booking, []BookingSeat, error)
	GetBookingByID(ctx context.Context, bookingId int) (*Booking, error)
	ListBookingsByUser(ctx context.Context, userId int) ([]Booking, error)
}

func NewService(db *gorm.DB, repo Repository, movieRepo movies.Repository, theaterRepo theatres.Repository) Service {
	return &service{
		db:          db,
		repo:        repo,
		movieRepo:   movieRepo,
		theaterRepo: theaterRepo,
	}
}
func (s *service) CreateBooking(ctx context.Context, createReq CreateBookingRequest) (*Booking, []BookingSeat, error) {
	seats, err := s.theaterRepo.GetSeatsByIds(ctx, createReq.SeatIDs)
	if err != nil {
		return nil, nil, err
	}

	if len(seats) == 0 {
		return nil, nil, fmt.Errorf("no valid seats found for the provided seat IDs")
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := s.checkSeatAvailability(ctx, tx, createReq.ShowtimeID, createReq.SeatIDs); err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	totalAmount := 0.0
	for _, seat := range seats {
		totalAmount += seat.SeatCategoryPrice
	}
	booking := &Booking{
		UserID:        uint(createReq.UserID),
		ShowtimeID:    uint(createReq.ShowtimeID),
		ScreenID:      createReq.ScreenID,
		BookingDate:   time.Now(),
		TotalAmount:   totalAmount,
		PaymentStatus: "Pending",
	}
	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	var bookingSeats []BookingSeat
	for _, seat := range seats {
		bookingSeat := BookingSeat{
			BookingID: booking.BookingID,
			SeatID:    seat.ID,
		}
		bookingSeats = append(bookingSeats, bookingSeat)
	}

	if err := tx.Create(&bookingSeats).Error; err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}

	return booking, bookingSeats, nil
}

func (s *service) checkSeatAvailability(ctx context.Context, tx *gorm.DB, showtimeID int, seatIDs []int) error {
	var existingBookings []BookingSeat
	err := tx.Table("booking_seats").
		Joins("JOIN bookings ON bookings.booking_id = booking_seats.booking_id").
		Where("bookings.showtime_id = ? AND booking_seats.seat_id IN ?", showtimeID, seatIDs).
		Find(&existingBookings).Error

	if err != nil {
		return err
	}
	if len(existingBookings) > 0 {
		return fmt.Errorf("one or more of the requested seats are already booked")
	}

	return nil
}

func (s *service) GetBookingByID(ctx context.Context, bookingId int) (*Booking, error) {
	bookings, err := s.repo.GetBookingByID(ctx, bookingId)
	if err != nil {
		return nil, err
	}
	return bookings, err
}

func (s *service) ListBookingsByUser(ctx context.Context, userId int) ([]Booking, error) {
	bookings, err := s.repo.ListBookingsByUser(ctx, userId)
	if len(bookings) < 1 {
		return nil, fmt.Errorf("no bookings found with user id %d", userId)
	}
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
