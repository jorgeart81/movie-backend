package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jorgeart81/movie-backend/cmd/config"
	"github.com/jorgeart81/movie-backend/cmd/router/controllers"
)

func MainRouter(app *config.Application) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	controllers.Init(app)

	mux.Route("/api", func(mux chi.Router) {
		mux.Get("/", controllers.Home)

		mux.Route("/movies", func(mux chi.Router) {
			mux.Get("/", controllers.Movies)
		})

	})

	return mux
}
