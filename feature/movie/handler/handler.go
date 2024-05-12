package handler

import (
	"context"
	"encoding/json"
	"github.com/RianIhsan/go-chi-crud/feature/movie"
	"github.com/RianIhsan/go-chi-crud/model"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *movieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	res, err := h.movieService.GetMovies(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *movieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	// Parse movie ID from URL path
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Call service to get movie by ID
	ctx := context.Background()
	movie, err := h.movieService.GetMovie(ctx, movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the movie data
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (h *movieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Parse movie ID from URL path
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Decode JSON request body to MovieRequest
	var req model.MovieRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Call service to update the movie
	ctx := context.Background()
	res, err := h.movieService.UpdateMovie(ctx, movieID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *movieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// Parse movie ID from URL path
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Call service to delete the movie
	ctx := context.Background()
	res, err := h.movieService.DeleteMovie(ctx, movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
