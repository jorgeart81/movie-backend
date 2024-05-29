package controllers

import (
	"fmt"
	"net/http"
)

func (c *ApiController) Movies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Movies from %s", c.Domain)
}

func (c *ApiController) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := c.Repository.AllMovies()
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = c.writeJSON(w, http.StatusOK, movies)
}
