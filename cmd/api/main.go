package main

import (
	"backend/cmd/config"
	"backend/cmd/router"
	"fmt"
	"log"
	"net/http"
)

type application struct {
	Domain string
}

func main() {
	var app application

	router := router.MainRouter()
	envs, _ := config.Envs()
	// TODO: set application config

	// TODO: read from command line

	// TODO: connect to the database

	app.Domain = envs.Domain

	// start a web server
	log.Println("Starting application on port", envs.ApiPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", envs.ApiPort), router)
	if err != nil {
		log.Fatal(err)
	}
}
