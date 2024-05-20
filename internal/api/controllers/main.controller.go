package controllers

import (
	"encoding/json"
	"fmt"
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

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
