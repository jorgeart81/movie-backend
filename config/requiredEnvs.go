package config

type environmentKeys struct {
	API_PORT          string
	API_HOST          string
	DOMAIN            string
	CORS_ALLOW_ORIGIN string
}

var envKeys = environmentKeys{
	API_PORT:          "API_PORT",
	API_HOST:          "API_HOST",
	DOMAIN:            "DOMAIN",
	CORS_ALLOW_ORIGIN: "CORS_ALLOW_ORIGIN",
}

var requiredEnvs = []string{
	envKeys.API_PORT,
	envKeys.API_HOST,
	envKeys.DOMAIN,
	envKeys.CORS_ALLOW_ORIGIN,
}
