package movies

import (
	"context"
	"fmt"
	"time"

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
func (h *GrpcHandler) DeleteMovie(ctx context.Context, req *movie_booking.DeleteMovieRequest) (*movie_booking.DeleteMovieResponse, error) {
	err := h.svc.DeleteMovie(ctx, int(req.MovieId))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *GrpcHandler) GetMovieDetails(ctx context.Context, req *movie_booking.GetMovieDetailsRequest) (*movie_booking.GetMovieDetailsResponse, error) {
	movie, err := h.svc.GetMovieDetails(ctx, int(req.MovieId))
	if err != nil {
		return nil, err
	}
	return &movie_booking.GetMovieDetailsResponse{
		Movie: &movie_booking.Movie{
			MovieId:     req.MovieId,
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    int32(movie.Duration),
			Genre:       movie.Genre,
			ReleaseDate: movie.ReleaseDate.String(),
			Rating:      float32(movie.Rating),
			Language:    movie.Language,
		},
	}, nil
}
func (h *GrpcHandler) UpdateMovie(ctx context.Context, req *movie_booking.UpdateMovieRequest) (*movie_booking.UpdateMovieResponse, error) {
	var date *time.Time
	var err error
	if req.ReleaseDate != "" {
		date, err = utils.ParseDateString(req.ReleaseDate)
		if err != nil {
			return nil, fmt.Errorf("invalid date format for ReleaseDate: %w", err)
		}
	}

	movie := Movie{
		Title:       req.Title,
		Description: req.Description,
		Duration:    int(req.Duration),
		Genre:       req.Genre,
		Rating:      float64(req.Rating),
		Language:    req.Language,
	}
	if date != nil {
		movie.ReleaseDate = *date
	}

	err = h.svc.UpdateMovie(ctx, movie, int(req.MovieId))
	if err != nil {
		return nil, err
	}

	return &movie_booking.UpdateMovieResponse{
		Message: "Movie updated successfully",
	}, nil
}

func (h *GrpcHandler) ListMovies(ctx context.Context, req *movie_booking.ListMoviesRequest) (*movie_booking.ListMoviesResponse, error) {
	response, err := h.svc.ListMovies(ctx)
	if err != nil {
		return nil, err
	}

	var grpcMovies []*movie_booking.Movie
	for _, m := range response {
		grpcMovie := &movie_booking.Movie{
			Title:       m.Title,
			Description: m.Description,
			Duration:    int32(m.Duration),
			Genre:       m.Genre,
			ReleaseDate: m.ReleaseDate.String(),
			Rating:      float32(m.Rating),
			Language:    m.Language,
		}
		grpcMovies = append(grpcMovies, grpcMovie)
	}

	return &movie_booking.ListMoviesResponse{
		Movies: grpcMovies,
	}, nil
}
