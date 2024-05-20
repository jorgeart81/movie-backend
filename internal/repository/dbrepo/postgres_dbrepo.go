package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/jorgeart81/movie-backend/internal/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

// Connection implements repository.DatabaseRepo.
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

const dbTimeout = time.Second * 3

// AllMovies implements repository.DatabaseRepo.
func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
					id, title, release_date, runtime, 
					mpaa_rating, description, coalesce(image, ''),
					created_at, updated_at
		FROM
				movies
		ORDER BY
				title
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}