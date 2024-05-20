package controllers

import (
	"encoding/json"
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

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
