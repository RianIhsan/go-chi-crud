package routes

import (
	"github.com/RianIhsan/go-chi-crud/feature/movie"
	"github.com/RianIhsan/go-chi-crud/feature/movie/handler"
	"github.com/go-chi/chi/v5"
)

func MovieRoutes(r chi.Router, movieService movie.MovieServiceInterface) {
	movieHandler := handler.NewMovieHandler(movieService)
	r.Post("/v1/movies", movieHandler.CreateMovie)
}
