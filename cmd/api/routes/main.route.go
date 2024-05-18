package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jorgeart81/movie-backend/cmd/api/middlewares"
	"github.com/jorgeart81/movie-backend/cmd/api/routes/controllers"
	"github.com/jorgeart81/movie-backend/internal/models"
)

func MainRouter(app *models.Application) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.EnableCORS)

	controllers.Init(app)

	mux.Route("/api", func(mux chi.Router) {
		mux.Get("/", controllers.Home)

		mux.Route("/movies", func(mux chi.Router) {
			mux.Get("/", controllers.AllMovies)
		})

	})

	return mux
}
