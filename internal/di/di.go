package di

import (
	"log"

	"github.com/aparnasukesh/movies-booking-svc/config"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/booking"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/movies"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/theatres"
	"github.com/aparnasukesh/movies-booking-svc/internal/boot"
	grpclient "github.com/aparnasukesh/movies-booking-svc/pkg/grpClient"
	redis "github.com/aparnasukesh/movies-booking-svc/pkg/redis"
	sql "github.com/aparnasukesh/movies-booking-svc/pkg/sql"
)

func InitResources(cfg config.Config) (func() error, error) {

	// Db initialization
	db, err := sql.NewSql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Redis initialization
	redisClient, err := redis.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// // Movie Module Initialization
	movieRepo := movies.NewRepository(db)
	movieService := movies.NewService(movieRepo, redisClient)
	movieGrpcHandler := movies.NewGrpcHandler(movieService)

	// Theatres Module initialization
	theaterRepo := theatres.NewRepository(db)
	service := theatres.NewService(theaterRepo, movieRepo)
	theatresGrpcHandler := theatres.NewGrpcHandler(service)

	// Booking Module Initialization
	paymentSvcClient, err := grpclient.NewBookingPaymentServiceClient(cfg.GrpcPaymentPort)
	if err != nil {
		return nil, err
	}
	bookingRepo := booking.NewRepository(db)
	bookingService := booking.NewService(db, bookingRepo, movieRepo, theaterRepo, paymentSvcClient)
	bookingGrpcHandler := booking.NewGrpcHandler(bookingService)

	// Server initialization
	server, err := boot.NewGrpcServer(cfg, movieGrpcHandler, theatresGrpcHandler, bookingGrpcHandler)
	if err != nil {
		log.Fatal(err)
	}
	return server, nil
}
