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

	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         int
	DBHost         string
	SSLMode        string
	Timezone       string
	ConnectTimeOut string
}

func Envs() *Environment {
	// Load if .env file exists
	loadEnv()

	err := checkEnvVars(requiredEnvs)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &Environment{
		APIPort:         parseInt(os.Getenv(envKeys.API_PORT), "error parsing API_PORT"),
		APIHost:         os.Getenv(envKeys.API_HOST),
		Domain:          os.Getenv(envKeys.DOMAIN),
		CORSAllowOrigin: os.Getenv(envKeys.CORS_ALLOW_ORIGIN),

		DBUser:         os.Getenv(envKeys.POSTGRES_USER),
		DBPassword:     os.Getenv(envKeys.POSTGRES_PASSWORD),
		DBName:         os.Getenv(envKeys.POSTGRES_DB),
		DBPort:         parseInt(os.Getenv(envKeys.DB_PORT), "error parsing DB_PORT"),
		DBHost:         os.Getenv(envKeys.DB_HOST),
		SSLMode:        os.Getenv(envKeys.SSLMODE),
		Timezone:       os.Getenv(envKeys.TIMEZONE),
		ConnectTimeOut: os.Getenv(envKeys.CONNECT_TIMEOUT),
	}
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
