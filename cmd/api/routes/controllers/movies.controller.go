package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jorgeart81/movie-backend/internal/models"
)

func Movies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Movies from %s", controllerStruct.app.Domain)
}

func AllMovies(w http.ResponseWriter, r *http.Request) {
	rd, _ := time.Parse("2006-01-02", "1986-03-07")
	var movies = []models.Movie{
		{
			ID:          1,
			Title:       "Highlander",
			ReleaseDate: rd,
			MPAARating:  "R",
			Runtime:     116,
			Description: "A very nice movie",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "Raiders of the Lost Ark",
			ReleaseDate: rd,
			MPAARating:  "PG",
			Runtime:     115,
			Description: "Another very nice movie",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
