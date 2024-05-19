package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jorgeart81/movie-backend/cmd/api/routes"
	"github.com/jorgeart81/movie-backend/config"
	"github.com/jorgeart81/movie-backend/internal/models"
)

func main() {
	var app models.Application
	envs := config.Envs()

	flag.StringVar(&app.Domain, "domain", envs.Domain, "domain")
	flag.Parse()

	router := routes.MainRouter(&app)

	// TODO: set application config

	// TODO: read from command line

	// TODO: connect to the database

	// start a web server
	log.Println("Starting application on port", envs.ApiPort)
	addr := fmt.Sprintf("%s:%d", envs.ApiHost, envs.ApiPort)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
