package models

type environmentKeys struct {
	API_PORT          string
	API_HOST          string
	DOMAIN            string
	CORS_ALLOW_ORIGIN string
}

var EnvKeys = environmentKeys{
	API_PORT:          "API_PORT",
	API_HOST:          "API_HOST",
	DOMAIN:            "DOMAIN",
	CORS_ALLOW_ORIGIN: "CORS_ALLOW_ORIGIN",
}
