package repository

import (
	"context"
	"database/sql"
	"github.com/RianIhsan/go-chi-crud/feature/movie"
	"github.com/RianIhsan/go-chi-crud/model"
)

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) movie.MovieRepositoryInterface {
	return &movieRepository{db: db}
}

var (
	InsertMovieQuery = `INSERT INTO mst_movie (title, author ,description) VALUES ($1, $2, $3)`
)

func (m *movieRepository) InsertMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error) {
	_, err := m.db.QueryContext(ctx, InsertMovieQuery, req.Title, req.Author, req.Description)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: "Success Insert Movie",
	}, nil
}
