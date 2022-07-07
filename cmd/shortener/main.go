package main

import (
	"net/http"
	"time"

	"github.com/ArtemBonda/shortener/internal/app"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.RootAcceptURL)

	server := &http.Server{
		Addr:        "localhost:8080",
		Handler:     mux,
		ReadTimeout: 1 * time.Second,
	}
	server.ListenAndServe()
}
