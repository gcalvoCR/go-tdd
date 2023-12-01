package main

import (
	"net/http"
	"time"

	"github.com/gcalvocr/go-tdd/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", routes.HandleHealthCheck)

	s := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
