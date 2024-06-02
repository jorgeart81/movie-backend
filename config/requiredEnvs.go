package config

type environmentKeys struct {
	API_PORT          string
	API_HOST          string
	CORS_ALLOW_ORIGIN string
	DOMAIN            string
	COOKIE_DOMAIN     string
	JWT_SECRET        string
	JWT_ISSUER        string
	JWT_AUDIENCE      string

	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	DB_PORT           string
	DB_HOST           string
	SSLMODE           string
	TIMEZONE          string
	CONNECT_TIMEOUT   string
}

var envKeys = environmentKeys{
	API_PORT:          "API_PORT",
	API_HOST:          "API_HOST",
	CORS_ALLOW_ORIGIN: "CORS_ALLOW_ORIGIN",
	DOMAIN:            "DOMAIN",
	COOKIE_DOMAIN:     "COOKIE_DOMAIN",
	JWT_SECRET:        "JWT_SECRET",
	JWT_ISSUER:        "JWT_ISSUER",
	JWT_AUDIENCE:      "JWT_AUDIENCE",

	POSTGRES_USER:     "POSTGRES_USER",
	POSTGRES_PASSWORD: "POSTGRES_PASSWORD",
	POSTGRES_DB:       "POSTGRES_DB",
	DB_PORT:           "DB_PORT",
	DB_HOST:           "DB_HOST",
	SSLMODE:           "SSLMODE",
	TIMEZONE:          "TIMEZONE",
	CONNECT_TIMEOUT:   "CONNECT_TIMEOUT",
}

var requiredEnvs = []string{
	envKeys.API_PORT,
	envKeys.API_HOST,
	envKeys.CORS_ALLOW_ORIGIN,
	envKeys.DOMAIN,
	envKeys.JWT_SECRET,
	envKeys.JWT_ISSUER,
	envKeys.JWT_AUDIENCE,

	envKeys.POSTGRES_USER,
	envKeys.POSTGRES_PASSWORD,
	envKeys.POSTGRES_DB,
	envKeys.DB_PORT,
	envKeys.DB_HOST,
	envKeys.SSLMODE,
	envKeys.TIMEZONE,
	envKeys.CONNECT_TIMEOUT,
}
