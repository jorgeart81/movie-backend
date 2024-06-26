package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jorgeart81/movie-backend/internal/api/controllers"
)

func (s *server) initRouter() {
	s.router = chi.NewRouter()
	controller := controllers.ApiController{
		Domain:     s.Domain,
		Repository: s.DB,
		Auth:       &s.auth,
	}

	s.router.Use(middleware.Recoverer)
	s.router.Use(s.enableCORS)

	s.routes(&controller)
	s.protectedRoutes(&controller)
}

func (s *server) routes(c *controllers.ApiController) {
	s.router.Route("/api", func(mux chi.Router) {
		mux.Get("/", c.Home)

		mux.Route("/authenticate", func(mux chi.Router) {
			mux.Post("/", c.Authenticate)
			mux.Get("/refresh", c.RefreshToken)
			mux.Get("/logout", c.Logout)
		})

		mux.Get("/movies/{id}", c.GetMovie)
	})
}

func (s *server) protectedRoutes(c *controllers.ApiController) {
	s.router.Route("/api/admin", func(mux chi.Router) {
		mux.Use(s.authRequired)

		mux.Get("/movies", c.AllMovies)
		mux.Get("/movies/{id}", c.MovieForEdit)

	})
}
