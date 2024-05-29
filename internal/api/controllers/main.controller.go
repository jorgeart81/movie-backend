package controllers

import (
	"net/http"

	"github.com/jorgeart81/movie-backend/internal/repository"
)

type ApiController struct {
	Domain     string
	Repository repository.DatabaseRepo
}

func (c *ApiController) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	_ = c.writeJSON(w, http.StatusOK, payload)
}
