package movies

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type service struct {
	repo        Repository
	redisClient *redis.Client
}

type Service interface {
	RegisterMovie(ctx context.Context, movie Movie) (int, error)
	UpdateMovie(ctx context.Context, movie Movie, movieId int) error
	ListMovies(ctx context.Context) ([]Movie, error)
	GetMovieDetailsByID(ctx context.Context, movieId int) (*Movie, error)
	DeleteMovie(ctx context.Context, movieId int) error
	GetMoviesByLanguage(ctx context.Context, language string) ([]Movie, error)
	GetMoviesByGenre(ctx context.Context, genre string) ([]Movie, error)
	GetMovieByName(ctx context.Context, name string) (*Movie, error)
	GetMovieByNameAndLanguage(ctx context.Context, name, language string) (*Movie, error)
	// Redis
	GetFromCache(ctx context.Context, cacheKey string, result interface{}) error
	SetToCache(ctx context.Context, cacheKey string, data interface{}, expiration time.Duration) error
}

func NewService(repo Repository, redisClient *redis.Client) Service {
	return &service{
		repo:        repo,
		redisClient: redisClient,
	}
}

func (s *service) GetFromCache(ctx context.Context, cacheKey string, result interface{}) error {
	val, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cache data: %w", err)
	}

	return nil
}

func (s *service) SetToCache(ctx context.Context, cacheKey string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data for cache: %w", err)
	}

	err = s.redisClient.Set(ctx, cacheKey, jsonData, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache: %w", err)
	}

	return nil
}

// Movies
func (s *service) RegisterMovie(ctx context.Context, movie Movie) (int, error) {
	res, err := s.repo.FindMovieByNameAndLanguage(ctx, movie)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
	}
	if res != nil && err == nil {
		return 0, errors.New("this movie already exist")
	}
	movieId, err := s.repo.CreateMovie(ctx, movie)
	if err != nil {
		return 0, err
	}
	return movieId, nil
}

func (s *service) DeleteMovie(ctx context.Context, movieId int) error {
	err := s.repo.DeleteMovie(ctx, movieId)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("movie:%d", movieId)
	err = s.redisClient.Del(ctx, cacheKey).Err()
	if err != nil {
		fmt.Printf("Failed to invalidate cache for movie %d: %v\n", movieId, err)
	}

	return nil
}

func (s *service) GetMovieDetailsByID(ctx context.Context, movieId int) (*Movie, error) {
	cacheKey := fmt.Sprintf("movie:%d", movieId)
	var cachedMovie Movie
	err := s.GetFromCache(ctx, cacheKey, &cachedMovie)
	if err == nil {
		fmt.Println("Movie found in cache")
		return &cachedMovie, nil
	} else if err != redis.Nil {
		fmt.Printf("Redis error: %v\n", err)
	}
	movie, err := s.repo.GetMovieDetailsById(ctx, movieId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("movie not found with the id %d", movieId)
	}
	err = s.SetToCache(ctx, cacheKey, movie, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache movie details: %v\n", err)
	}
	return movie, nil
}

func (s *service) GetMovieByName(ctx context.Context, name string) (*Movie, error) {
	cacheKey := fmt.Sprintf("movie:%s", name)
	var cachedMovie Movie
	err := s.GetFromCache(ctx, cacheKey, &cachedMovie)
	if err == nil {
		fmt.Println("Movie found in cache")
		return &cachedMovie, nil
	} else if err != redis.Nil {
		fmt.Printf("Redis error: %v\n", err)
	}
	movie, err := s.repo.GetMovieByName(ctx, name)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("movie not found with the id %d", name)
	}
	err = s.SetToCache(ctx, cacheKey, movie, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache movie details: %v\n", err)
	}
	return movie, nil
}
func (s *service) GetMovieByNameAndLanguage(ctx context.Context, name, language string) (*Movie, error) {
	cacheKey := fmt.Sprintf("movie:%s:%s", name, language)
	var cachedMovie Movie
	err := s.GetFromCache(ctx, cacheKey, &cachedMovie)
	if err == nil {
		fmt.Println("Movie found in cache")
		return &cachedMovie, nil
	} else if err != redis.Nil {
		fmt.Printf("Redis error: %v\n", err)
	}
	movie, err := s.repo.GetMovieByNameAndLanguage(ctx, name, language)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("movie not found with the name '%s' and language '%s'", name, language)
	}
	err = s.SetToCache(ctx, cacheKey, movie, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache movie details: %v\n", err)
	}

	return movie, nil
}

func (s *service) ListMovies(ctx context.Context) ([]Movie, error) {
	movies, err := s.repo.GetMovies(ctx)
	if err != nil {
		return nil, err
	}
	if len(movies) < 1 {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}

func (s *service) UpdateMovie(ctx context.Context, movie Movie, movieId int) error {
	err := s.repo.UpdateMovie(ctx, movie, movieId)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("movie:%d", movieId)
	err = s.redisClient.Del(ctx, cacheKey).Err()
	if err != nil {
		fmt.Printf("Failed to invalidate cache for movie %d: %v\n", movieId, err)
	}

	return nil
}

func (s *service) GetMoviesByGenre(ctx context.Context, genre string) ([]Movie, error) {
	movies, err := s.repo.GetMoviesByGenre(ctx, genre)
	if err != nil {
		return nil, err
	}
	if len(movies) < 1 {
		return nil, fmt.Errorf("no movies found in this genre %s", genre)
	}
	return movies, nil
}

func (s *service) GetMoviesByLanguage(ctx context.Context, language string) ([]Movie, error) {
	movies, err := s.repo.GetMoviesByLanguage(ctx, language)
	if err != nil {
		return nil, err
	}
	if len(movies) < 1 {
		return nil, fmt.Errorf("no movies found in this language  %s", language)
	}
	return movies, nil
}
