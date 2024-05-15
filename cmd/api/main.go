package main

import (
	"backend/cmd/config"
	"fmt"
	"log"
	"net/http"
)

type application struct {
	Domain string
}

func main() {
	var app application
	envs, _ := config.Envs()
	// TODO: set application config

	// TODO: read from command line

	// TODO: connect to the database

	app.Domain = envs.Domain

	http.HandleFunc("/", Hello)

	// start a web server
	log.Println("Starting application on port", envs.ApiPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", envs.ApiPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}
