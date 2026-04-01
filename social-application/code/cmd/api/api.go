package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Swaniket/social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

// As chi internally implements the http.Handler interface, we can use it directly as the Handler for our http.Server.
// Previously we had to use *chi.Mux

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// // Using a middleware - This will recover from any panics and return a 500 Internal Server Error status code. It also logs the panic message and stack trace to the console.
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Logger) // logger

	r.Use(middleware.RequestID) // request ID
	r.Use(middleware.RealIP)    // real IP address of the client
	r.Use(middleware.Logger)    // logger
	r.Use(middleware.Recoverer) // recover from panics

	// Set a timeout value on the request context (ctx), that will signal through ctx.Done()
	// that the request has timed out and further processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// r.Get("/health", app.healthCheckHandler)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Timeout for writing the response - 30 seconds
		ReadTimeout:  time.Second * 10, // Timeout for reading the request - 10 seconds
		IdleTimeout:  time.Minute,      // Timeout for idle connections - 1 minute
	}

	log.Printf("Server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
