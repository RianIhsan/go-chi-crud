package handler

import (
	"context"
	"encoding/json"
	"github.com/RianIhsan/go-chi-crud/feature/movie"
	"github.com/RianIhsan/go-chi-crud/model"
	"net/http"
)

type movieHandler struct {
	movieService movie.MovieServiceInterface
}

func NewMovieHandler(movieService movie.MovieServiceInterface) movie.MovieHandlerInterface {
	return &movieHandler{movieService}
}

func (h *movieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var req model.MovieRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := context.Background()

	res, err := h.movieService.CreateMovie(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
