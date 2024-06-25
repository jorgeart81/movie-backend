package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jorgeart81/movie-backend/internal/models"
)

func (c *ApiController) Movies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Movies from %s", c.Domain)
}

func (c *ApiController) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := c.Repository.AllMovies()
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	_ = c.writeJSON(w, http.StatusOK, movies)
}

func (c *ApiController) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	movie, err := c.Repository.OneMovie(movieID)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	_ = c.writeJSON(w, http.StatusOK, movie)
}

func (c *ApiController) MovieForEdit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	movieID, err := strconv.Atoi(id)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	movie, genres, err := c.Repository.OneMovieForEdit(movieID)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	var payload = struct {
		Movie  *models.Movie   `json:"movie"`
		Genres []*models.Genre `json:"genres"`
	}{
		movie,
		genres,
	}

	_ = c.writeJSON(w, http.StatusOK, payload)

}
