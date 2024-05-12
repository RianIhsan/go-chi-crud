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

func (m *movieService) GetMovies(ctx context.Context) ([]model.MoviesResponse, error) {
	res, err := m.movieRepo.FindMovies(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *movieService) GetMovie(ctx context.Context, id int) (model.MoviesResponse, error) {
	res, err := m.movieRepo.FindMovie(ctx, id)
	if err != nil {
		return model.MoviesResponse{}, err
	}
	return res, nil
}

func (m *movieService) UpdateMovie(ctx context.Context, id int, req model.MovieRequest) (model.MovieResponse, error) {
	exisitngData, err := m.movieRepo.FindMovie(ctx, id)
	if err != nil {
		return model.MovieResponse{}, err
	}

	if req.Title == "" {
		req.Title = exisitngData.Title
	}
	if req.Author == "" {
		req.Author = exisitngData.Author
	}
	if req.Description == "" {
		req.Description = exisitngData.Description

	}
	res, err := m.movieRepo.UpdateMovie(ctx, id, req)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: res.Message,
	}, nil
}

func (m *movieService) DeleteMovie(ctx context.Context, id int) (model.MovieResponse, error) {
	res, err := m.movieRepo.DeleteMovie(ctx, id)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: res.Message,
	}, nil
}
