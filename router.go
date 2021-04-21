package main

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func registerRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ping", ping).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/ping", ping).Methods(http.MethodGet)
	return r
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong"))
}
