package main

import (
	"fmt"
	"log"

	"github.com/jorgeart81/movie-backend/config"
	"github.com/jorgeart81/movie-backend/internal/api"
)

func main() {
	envs := config.Envs()

	// TODO: set application config

	// TODO: read from command line
	// flag.StringVar(&app.Domain, "domain", envs.Domain, "domain")
	// flag.Parse()

	// TODO: connect to the database

	// start a web server
	addr := fmt.Sprintf("%s:%d", envs.APIHost, envs.APIPort)
	server := api.NewServer(envs)
	// log.Fatal(server.Listen(addr))
	if err := server.Listen(addr); err != nil {
		log.Fatal(err)
	}

}
