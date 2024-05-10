package movie

import (
	"context"
	"github.com/RianIhsan/go-chi-crud/model"
	"net/http"
)

type MovieRepositoryInterface interface {
	InsertMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error)
}

type MovieServiceInterface interface {
	CreateMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error)
}

type MovieHandlerInterface interface {
	CreateMovie(w http.ResponseWriter, r *http.Request)
}
