package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jorgeart81/movie-backend/config"
	"github.com/jorgeart81/movie-backend/internal/api/controllers"
	"github.com/jorgeart81/movie-backend/internal/repository"
)

// Server serves the public run API.
type Server struct {
	router          *chi.Mux
	CORSAllowOrigin string
	Domain          string
	DB              repository.DatabaseRepo
}

// NewServer returns
func NewServer(envs *config.Environment, db repository.DatabaseRepo) *Server {
	log.Println("Starting application on port", envs.APIPort)

	return &Server{
		CORSAllowOrigin: envs.CORSAllowOrigin,
		Domain:          envs.Domain,
		DB:              db,
	}
}

// Listen starts listening on the given address.
func (s *Server) Listen(addr string) error {
	s.initRouter()
	return http.ListenAndServe(addr, s.router)

}

func (s *Server) initRouter() {
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
