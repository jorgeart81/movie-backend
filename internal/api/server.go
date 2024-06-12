package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jorgeart81/movie-backend/config"
	"github.com/jorgeart81/movie-backend/internal/api/auth"
	"github.com/jorgeart81/movie-backend/internal/repository"
)

// server serves the public run API.
type server struct {
	router          *chi.Mux
	CORSAllowOrigin string
	Domain          string
	DB              repository.DatabaseRepo
	auth            auth.Auth
}

// NewServer returns
func NewServer(envs *config.Environment, db repository.DatabaseRepo) *server {
	log.Println("Starting application on port", envs.APIPort)

	return &server{
		CORSAllowOrigin: envs.CORSAllowOrigin,
		Domain:          envs.Domain,
		DB:              db,
		auth: auth.Auth{
			Issuer:        envs.JWTIssuer,
			Audience:      envs.JWTAudience,
			Secret:        envs.JWTSecret,
			TokenExpiry:   time.Minute * 15,
			RefreshExpiry: time.Hour * 24,
			CookiePath:    "/",
			CookieName:    "Authorization",
			CookieDomain:  envs.CookieDomain,
		},
	}
}

// Listen starts listening on the given address.
func (s *server) Listen(addr string) error {
	s.initRouter()
	return http.ListenAndServe(addr, s.router)
}
