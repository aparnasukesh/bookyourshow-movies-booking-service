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
