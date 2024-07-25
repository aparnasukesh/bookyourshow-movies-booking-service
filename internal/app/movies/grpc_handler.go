package movies

import (
	"context"

	"github.com/aparnasukesh/inter-communication/movie_booking"
	"github.com/aparnasukesh/movies-booking-svc/pkg/utils"
)

type GrpcHandler struct {
	svc Service
	movie_booking.UnimplementedMovieServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}

func (h *GrpcHandler) RegisterMovie(ctx context.Context, req *movie_booking.RegisterMovieRequest) (*movie_booking.RegisterMovieResponse, error) {
	date, err := utils.ParseDateString(req.ReleaseDate)
	if err != nil {
		return nil, err
	}
	movie := Movie{
		Title:       req.Title,
		Description: req.Description,
		Duration:    int(req.Duration),
		Genre:       req.Genre,
		ReleaseDate: *date,
		Rating:      float64(req.Rating),
		Language:    req.Language,
	}
	movieId, err := h.svc.RegisterMovie(ctx, movie)
	if err != nil {
		return &movie_booking.RegisterMovieResponse{
			MovieId: 0,
			Message: "failed to create movie",
		}, err
	}
	return &movie_booking.RegisterMovieResponse{
		MovieId: uint32(movieId),
		Message: "create movie successfull",
	}, nil
}
