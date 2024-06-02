package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jorgeart81/movie-backend/config"
	"github.com/jorgeart81/movie-backend/internal/api/controllers"
	"github.com/jorgeart81/movie-backend/internal/repository"
)

// server serves the public run API.
type server struct {
	router          *chi.Mux
	CORSAllowOrigin string
	Domain          string
	DB              repository.DatabaseRepo
	auth            Auth
}

// NewServer returns
func NewServer(envs *config.Environment, db repository.DatabaseRepo) *server {
	log.Println("Starting application on port", envs.APIPort)

	return &server{
		CORSAllowOrigin: envs.CORSAllowOrigin,
		Domain:          envs.Domain,
		DB:              db,
		auth: Auth{
			Issuer:        envs.JWTIssuer,
			Audience:      envs.JWTAudience,
			Secret:        envs.JWTSecret,
			TokenExpiry:   time.Minute * 15,
			RefreshExpiry: time.Hour * 24,
			CookiePath:    "/",
			CookieName:    "__Host-refresh_token",
			CookieDomain:  envs.CookieDomain,
		},
	}
}

// Listen starts listening on the given address.
func (s *server) Listen(addr string) error {
	s.initRouter()
	return http.ListenAndServe(addr, s.router)

}

func (s *server) initRouter() {
	s.router = chi.NewRouter()
	controller := controllers.ApiController{
		Domain:     s.Domain,
		Repository: s.DB,
	}

	s.router.Use(middleware.Recoverer)
	s.router.Use(s.enableCORS)

	s.router.Route("/api", func(mux chi.Router) {
		mux.Get("/", controller.Home)

		mux.Route("/movies", func(mux chi.Router) {
			mux.Get("/", controller.AllMovies)
		})
	})
}
