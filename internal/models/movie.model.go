package models

import (
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"releaseDate"`
	Runtime     int       `json:"runtime"`
	MPAARating  string    `json:"mpaaRating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Genres      []*Genre  `json:"genres,omitempty"`
	GenresArray []int     `json:"genresArray,omitempty"`
}

func NewMovie(m Movie) *Movie {
	// deployID := uuid.New()

	return &Movie{
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		MPAARating:  m.MPAARating,
		Runtime:     m.Runtime,
		Description: m.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
