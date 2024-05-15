package router

import (
	"backend/cmd/router/controllers"
	"net/http"
)

type routes struct {
	apiV1 *http.ServeMux
}

func MainRouter() *http.ServeMux {
	var routes routes
	routes.init()

	mux := http.NewServeMux()
	mux.Handle("/api/", routes.apiV1)

	return mux
}

func (r *routes) init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Hello)

	r.apiV1 = mux
}
