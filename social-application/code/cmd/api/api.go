package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Swaniket/social/cmd/api/docs" // This is required to generate swagger docs
	"github.com/Swaniket/social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr   string
	db     dbConfig
	env    string
	apiUrl string
	mail   mailConfig
}

type mailConfig struct {
	exp time.Duration
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

	// All of the V1 routes
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.CreatePostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)
				r.Get("/", app.GetPostByIDHandler)
				r.Patch("/", app.UpdatePostByIDHandler) // Update Post
				r.Delete("/", app.DeletePostHandler)    // Delete Post
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Put("/activate/{token}", app.activateUserHandler)

			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.userContextMiddleware)
				r.Get("/", app.GetUserHandler)
				r.Put("/follow", app.FollowUserHandler)     // Follow
				r.Put("/unfollow", app.UnfollowUserHandler) // Follow
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
		})

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/user", app.RegisterUserHandler)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	// Docs
	docs.SwaggerInfo.Title = "Social"
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Description = "Simple social media implementation in Go"

	docs.SwaggerInfo.Host = app.config.apiUrl
	docs.SwaggerInfo.BasePath = "/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Timeout for writing the response - 30 seconds
		ReadTimeout:  time.Second * 10, // Timeout for reading the request - 10 seconds
		IdleTimeout:  time.Minute,      // Timeout for idle connections - 1 minute
	}

	app.logger.Infow("Server has started at", "addr", app.config.addr, "env", app.config.env)

	return srv.ListenAndServe()
}
