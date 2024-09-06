package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", indexHandler)
	mux.HandleFunc("GET /api/data", apiDataHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some data for the API"
	fmt.Fprintln(w, data)
}
