package service

import (
	"context"
	"github.com/RianIhsan/go-chi-crud/feature/movie"
	"github.com/RianIhsan/go-chi-crud/model"
)

type movieService struct {
	movieRepo movie.MovieRepositoryInterface
}

func NewMovieService(movieRepo movie.MovieRepositoryInterface) movie.MovieServiceInterface {
	return &movieService{movieRepo: movieRepo}
}

func (m *movieService) CreateMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error) {
	res, err := m.movieRepo.InsertMovie(ctx, req)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: res.Message,
	}, nil
}
