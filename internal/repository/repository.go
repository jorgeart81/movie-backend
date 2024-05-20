package repository

import (
	"database/sql"

	"github.com/jorgeart81/movie-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
