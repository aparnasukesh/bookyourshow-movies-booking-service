package di

import (
	"log"

	"github.com/aparnasukesh/movies-booking-svc/config"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/movies"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/theatres"
	"github.com/aparnasukesh/movies-booking-svc/internal/boot"
	"github.com/aparnasukesh/movies-booking-svc/pkg/sql"
)

func InitResources(cfg config.Config) (func() error, error) {

	// Db initialization
	db, err := sql.NewSql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// // Movie Module Initialization
	movieRepo := movies.NewRepository(db)
	movieService := movies.NewService(movieRepo)
	movieGrpcHandler := movies.NewGrpcHandler(movieService)

	// Theatres Module initialization
	repo := theatres.NewRepository(db)
	service := theatres.NewService(repo, movieRepo)
	theatresGrpcHandler := theatres.NewGrpcHandler(service)

	// Server initialization
	server, err := boot.NewGrpcServer(cfg, movieGrpcHandler, theatresGrpcHandler)
	if err != nil {
		log.Fatal(err)
	}
	return server, nil
}
