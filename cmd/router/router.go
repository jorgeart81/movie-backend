package router

import (
	"backend/cmd/router/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MainRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/api", func(mux chi.Router) {

		mux.Route("/movies", func(mux chi.Router) {
			mux.Get("/", controllers.Hello)
		})

	})

	return mux
}
