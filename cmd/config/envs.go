package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	ApiPort int
	Domain  string
}

func Envs() (Environment, error) {
	var env Environment

	err := loadEnv()
	if err != nil {
		return env, err
	}

	env.ApiPort = parseInt(os.Getenv("API_PORT"), "error parsing CONNECT_TIMEOUT")
	env.Domain = os.Getenv("DOMAIN")

	return env, nil
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
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
