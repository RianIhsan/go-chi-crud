package movie

import (
	"context"
	"github.com/RianIhsan/go-chi-crud/model"
	"net/http"
)

type MovieRepositoryInterface interface {
	InsertMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error)
	FindMovies(ctx context.Context) ([]model.MoviesResponse, error)
	FindMovie(ctx context.Context, id int) (model.MoviesResponse, error)
	UpdateMovie(ctx context.Context, id int, req model.MovieRequest) (model.MovieResponse, error)
	DeleteMovie(ctx context.Context, id int) (model.MovieResponse, error)
}

type MovieServiceInterface interface {
	CreateMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error)
	GetMovies(ctx context.Context) ([]model.MoviesResponse, error)
	GetMovie(ctx context.Context, id int) (model.MoviesResponse, error)
	UpdateMovie(ctx context.Context, id int, req model.MovieRequest) (model.MovieResponse, error)
	DeleteMovie(ctx context.Context, id int) (model.MovieResponse, error)
}

type MovieHandlerInterface interface {
	CreateMovie(w http.ResponseWriter, r *http.Request)
	GetMovies(w http.ResponseWriter, r *http.Request)
	GetMovie(w http.ResponseWriter, r *http.Request)
	UpdateMovie(w http.ResponseWriter, r *http.Request)
	DeleteMovie(w http.ResponseWriter, r *http.Request)
}
