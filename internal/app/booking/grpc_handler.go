package booking

import (
	"context"

	"github.com/aparnasukesh/inter-communication/movie_booking"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcHandler struct {
	svc Service
	movie_booking.UnimplementedBookingServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) CreateBooking(ctx context.Context, req *movie_booking.CreateBookingRequest) (*movie_booking.CreateBookingResponse, error) {

	seatIds := make([]int, len(req.SeatIds))
	for i := 0; i < len(req.SeatIds); i++ {
		seatIds[i] = int(req.SeatIds[i])
	}
	booking, bookingseats, err := h.svc.CreateBooking(ctx, CreateBookingRequest{
		UserID:      int(req.UserId),
		ShowtimeID:  int(req.ShowtimeId),
		SeatIDs:     seatIds,
		TotalAmount: req.TotalAmount,
	})
	if err != nil {
		return nil, err
	}
	bookingSeats := []*movie_booking.BookingSeat{}

	for _, seat := range bookingseats {
		res := movie_booking.BookingSeat{
			BookingId: uint32(seat.BookingID),
			SeatId:    uint32(seat.SeatID),
		}
		bookingSeats = append(bookingSeats, &res)
	}

	return &movie_booking.CreateBookingResponse{
		Booking: &movie_booking.Booking{
			BookingId:     uint32(booking.BookingID),
			UserId:        uint32(booking.UserID),
			ShowtimeId:    uint32(booking.ShowtimeID),
			BookingDate:   timestamppb.New(booking.BookingDate),
			TotalAmount:   booking.TotalAmount,
			PaymentStatus: booking.PaymentStatus,
			BookingSeats:  bookingSeats,
		},
	}, nil
}
func (h *GrpcHandler) GetBookingByID(ctx context.Context, req *movie_booking.GetBookingByIDRequest) (*movie_booking.GetBookingByIDResponse, error) {
	bookings, err := h.svc.GetBookingByID(ctx, int(req.BookingId))
	if err != nil {
		return nil, err
	}

	seats := make([]*movie_booking.BookingSeat, len(bookings.BookingSeats))
	for i, seat := range bookings.BookingSeats {
		seats[i] = &movie_booking.BookingSeat{
			BookingId: uint32(seat.BookingID),
			SeatId:    uint32(seat.SeatID),
		}
	}

	return &movie_booking.GetBookingByIDResponse{
		Booking: &movie_booking.Booking{
			BookingId:     uint32(bookings.BookingID),
			UserId:        uint32(bookings.UserID),
			ShowtimeId:    uint32(bookings.ShowtimeID),
			BookingDate:   timestamppb.New(bookings.BookingDate),
			TotalAmount:   bookings.TotalAmount,
			PaymentStatus: bookings.PaymentStatus,
			BookingSeats:  seats,
		},
	}, nil
}

func (h *GrpcHandler) ListBookingsByUser(ctx context.Context, req *movie_booking.ListBookingsByUserRequest) (*movie_booking.ListBookingsByUserResponse, error) {
	bookings, err := h.svc.ListBookingsByUser(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	response := []*movie_booking.Booking{}
	for _, booking := range bookings {
		seats := make([]*movie_booking.BookingSeat, len(booking.BookingSeats))
		for i, seat := range booking.BookingSeats {
			seats[i] = &movie_booking.BookingSeat{
				BookingId: uint32(seat.BookingID),
				SeatId:    uint32(seat.SeatID),
			}
		}
		res := movie_booking.Booking{
			BookingId:     uint32(booking.BookingID),
			UserId:        uint32(booking.UserID),
			ShowtimeId:    uint32(booking.ShowtimeID),
			BookingDate:   timestamppb.New(booking.BookingDate),
			TotalAmount:   booking.TotalAmount,
			PaymentStatus: booking.PaymentStatus,
			BookingSeats:  seats,
		}
		response = append(response, &res)
	}
	return &movie_booking.ListBookingsByUserResponse{
		Bookings: response,
	}, nil
}
