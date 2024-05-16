package controllers

import (
	"fmt"
	"net/http"

	"github.com/jorgeart81/movie-backend/cmd/config"
)

type apiController struct {
	app *config.Application
}

var controllerStruct apiController

func Init(app *config.Application) {
	controllerStruct.app = app
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Home from %s", controllerStruct.app.Domain)
}
