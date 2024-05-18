package controllers

import (
	"fmt"
	"net/http"
)

func Movies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Movies from %s", controllerStruct.app.Domain)
}
