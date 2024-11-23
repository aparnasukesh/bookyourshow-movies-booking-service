package boot

import (
	"log"
	"net"

	"github.com/aparnasukesh/inter-communication/movie_booking"
	"github.com/aparnasukesh/movies-booking-svc/config"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/booking"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/movies"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/theatres"
	"google.golang.org/grpc"
)

func NewGrpcServer(config config.Config, movieGrpcHandler movies.GrpcHandler, theatresGrpcHandler theatres.GrpcHandler, bookingGrpcHandler booking.GrpcHandler) (func() error, error) {
	//lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	lis, err := net.Listen("tcp", "0.0.0.0:"+config.GrpcPort)

	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	movie_booking.RegisterMovieServiceServer(s, &movieGrpcHandler)
	movie_booking.RegisterTheatreServiceServer(s, &theatresGrpcHandler)
	movie_booking.RegisterBookingServiceServer(s, &bookingGrpcHandler)
	srv := func() error {
		log.Printf("gRPC server started on port %s", config.GrpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}
		return nil
	}
	return srv, nil
}
