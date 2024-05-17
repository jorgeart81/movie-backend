package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jorgeart81/movie-backend/api/router"
	"github.com/jorgeart81/movie-backend/config"
)

func main() {
	var app config.Application

	router := router.MainRouter(&app)
	envs, _ := config.Envs()

	app.Domain = envs.Domain

	// TODO: set application config

	// TODO: read from command line

	// TODO: connect to the database

	flag.StringVar(&app.Domain, "domain", envs.Domain, "domain")
	flag.Parse()

	// start a web server
	log.Println("Starting application on port", envs.ApiPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", envs.ApiPort), router)
	if err != nil {
		log.Fatal(err)
	}
}
