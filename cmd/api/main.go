package main

import (
	"fmt"
	"log"

	"github.com/jorgeart81/movie-backend/config"
	"github.com/jorgeart81/movie-backend/internal/api"
	"github.com/jorgeart81/movie-backend/internal/repository/dbrepo"
)

func main() {
	envs := config.Envs()

	// Connect to the database
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
		envs.DBHost, envs.DBPort, envs.DBUser, envs.DBPassword, envs.DBName, envs.SSLMode, envs.Timezone, envs.ConnectTimeOut)

	conn, err := api.ConnectToDB(dns)
	if err != nil {
		log.Fatal(err)
	}

	dbrepo := &dbrepo.PostgresDBRepo{DB: conn}
	addr := fmt.Sprintf("%s:%d", envs.APIHost, envs.APIPort)

	// Start server
	server := api.NewServer(envs, dbrepo)
	// log.Fatal(server.Listen(addr))
	if err := server.Listen(addr); err != nil {
		log.Fatal(err)
	}

	defer server.DB.Connection().Close()
}
