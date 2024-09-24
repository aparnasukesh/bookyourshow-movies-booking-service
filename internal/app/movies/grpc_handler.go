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

func (h *GrpcHandler) GetMoviesByLanguage(ctx context.Context, req *movie_booking.GetMoviesByLanguageRequest) (*movie_booking.GetMoviesByLanguageResponse, error) {
	var movies []Movie
	movies, err := h.svc.GetMoviesByLanguage(ctx, req.Language)
	if err != nil {
		return nil, err
	}
	response := []*movie_booking.Movie{}
	for _, movie := range movies {
		res := &movie_booking.Movie{
			MovieId:     uint32(movie.ID),
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    int32(movie.Duration),
			Genre:       movie.Genre,
			ReleaseDate: movie.ReleaseDate.String(),
			Rating:      float32(movie.Rating),
			Language:    movie.Language,
		}
		response = append(response, res)
	}
	return &movie_booking.GetMoviesByLanguageResponse{
		Movie: response,
	}, nil
}

func (h *GrpcHandler) GetMoviesByGenre(ctx context.Context, req *movie_booking.GetMoviesByGenreRequest) (*movie_booking.GetMoviesByGenreResponse, error) {
	var movies []Movie
	movies, err := h.svc.GetMoviesByGenre(ctx, req.Genre)
	if err != nil {
		return nil, err
	}
	response := []*movie_booking.Movie{}
	for _, movie := range movies {
		res := &movie_booking.Movie{
			MovieId:     uint32(movie.ID),
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    int32(movie.Duration),
			Genre:       movie.Genre,
			ReleaseDate: movie.ReleaseDate.String(),
			Rating:      float32(movie.Rating),
			Language:    movie.Language,
		}
		response = append(response, res)
	}
	return &movie_booking.GetMoviesByGenreResponse{
		Movie: response,
	}, nil
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

func (h *GrpcHandler) GetMovieDetailsByID(ctx context.Context, req *movie_booking.GetMovieDetailsRequest) (*movie_booking.GetMovieDetailsResponse, error) {
	movie, err := h.svc.GetMovieDetailsByID(ctx, int(req.MovieId))
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

func (h *GrpcHandler) GetMovieByName(ctx context.Context, req *movie_booking.GetMovieByNameRequest) (*movie_booking.GetMovieByNameResponse, error) {
	movie, err := h.svc.GetMovieByName(ctx, req.MovieName)
	if err != nil {
		return nil, err
	}
	response := &movie_booking.Movie{
		MovieId:     uint32(movie.ID),
		Title:       movie.Title,
		Description: movie.Description,
		Duration:    int32(movie.Duration),
		Genre:       movie.Genre,
		ReleaseDate: movie.ReleaseDate.String(),
		Rating:      float32(movie.Rating),
		Language:    movie.Language,
	}
	return &movie_booking.GetMovieByNameResponse{
		Movie: response,
	}, nil
}
