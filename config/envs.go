package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/jorgeart81/movie-backend/internal/models"
)

type Environment struct {
	ApiPort         int
	ApiHost         string
	Domain          string
	CORSAllowOrigin string
}

func Envs() Environment {
	var requiredEnvs = []string{
		models.EnvKeys.API_PORT,
		models.EnvKeys.API_HOST,
		models.EnvKeys.DOMAIN,
		models.EnvKeys.CORS_ALLOW_ORIGIN,
	}

	loadEnv()
	err := checkEnvVars(requiredEnvs)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var env Environment
	env.ApiPort = parseInt(os.Getenv(models.EnvKeys.API_PORT), "error parsing API_PORT")
	env.ApiHost = os.Getenv(models.EnvKeys.API_HOST)
	env.Domain = os.Getenv(models.EnvKeys.DOMAIN)
	env.CORSAllowOrigin = os.Getenv(models.EnvKeys.CORS_ALLOW_ORIGIN)

	return env
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func checkEnvVars(vars []string) error {
	for _, v := range vars {
		if _, exists := os.LookupEnv(v); !exists {
			return fmt.Errorf("missing required environment variable: %s", v)
		}
	}
	return nil
}

func parseInt(value string, errorMessage string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(errorMessage+":", err)
		return 0
	}
	return intValue
}
