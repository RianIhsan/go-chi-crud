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
	InsertMovieQuery  = `INSERT INTO mst_movie (title, author ,description) VALUES ($1, $2, $3)`
	GetMoviesQuery    = `SELECT id, title, author, description, created_at, updated_at FROM mst_movie`
	GetMovieByIDQuery = `SELECT id, title, author, description, created_at, updated_at FROM mst_movie WHERE id = $1`
	UpdateMovieQuery  = `UPDATE mst_movie SET title = $1, author = $2, description = $3 WHERE id = $4`
	DeleteMovieQuery  = `DELETE FROM mst_movie WHERE id = $1`
)

func (m *movieRepository) InsertMovie(ctx context.Context, req model.MovieRequest) (model.MovieResponse, error) {
	_, err := m.db.QueryContext(ctx, InsertMovieQuery, req.Title, req.Author, req.Description)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: "success inset movie",
	}, nil
}

func (m *movieRepository) FindMovies(ctx context.Context) ([]model.MoviesResponse, error) {
	rows, err := m.db.QueryContext(ctx, GetMoviesQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var movies []model.MoviesResponse
	for rows.Next() {
		var movie model.MoviesResponse
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Author, &movie.Description, &movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (m *movieRepository) FindMovie(ctx context.Context, id int) (model.MoviesResponse, error) {
	var movie model.MoviesResponse
	err := m.db.QueryRowContext(ctx, GetMovieByIDQuery, id).Scan(&movie.Id, &movie.Title, &movie.Author, &movie.Description, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return model.MoviesResponse{}, err
	}
	return movie, nil
}

func (m *movieRepository) UpdateMovie(ctx context.Context, id int, req model.MovieRequest) (model.MovieResponse, error) {
	_, err := m.db.QueryContext(ctx, UpdateMovieQuery, req.Title, req.Author, req.Description, id)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: "success update movie",
	}, nil
}

func (m *movieRepository) DeleteMovie(ctx context.Context, id int) (model.MovieResponse, error) {
	_, err := m.db.ExecContext(ctx, DeleteMovieQuery, id)
	if err != nil {
		return model.MovieResponse{}, err
	}
	return model.MovieResponse{
		Message: "success delete movie",
	}, nil
}
