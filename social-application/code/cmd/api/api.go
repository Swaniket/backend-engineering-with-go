package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mountRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Timeout for writing the response
		ReadTimeout:  time.Second * 10, // Timeout for reading the request
	}

	log.Printf("Server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
