package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	APIPort         int
	APIHost         string
	Domain          string
	CORSAllowOrigin string
}

func Envs() *Environment {

	loadEnv()
	err := checkEnvVars(requiredEnvs)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var env Environment
	env.APIPort = parseInt(os.Getenv(envKeys.API_PORT), "error parsing API_PORT")
	env.APIHost = os.Getenv(envKeys.API_HOST)
	env.Domain = os.Getenv(envKeys.DOMAIN)
	env.CORSAllowOrigin = os.Getenv(envKeys.CORS_ALLOW_ORIGIN)

	return &env
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
