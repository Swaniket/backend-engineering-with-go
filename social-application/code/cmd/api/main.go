package main

import (
	"os"

	"github.com/Swaniket/social/internal/db"
	"github.com/Swaniket/social/internal/env"
	"github.com/Swaniket/social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1" // Pre-release

// @security					ApiKeyAuth
// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Enter your bearer token in the format "Bearer <token>"
func main() {
	// Having "postgres://user:adminpassword@localhost/social?sslmode=disable" - as the default
	// We are going to run postgres with docker for local, and for that we don't need DB_ADDR env variable
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiUrl: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	// Logger Setup
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal(err) // Don't want to bootup the server if the db connection fails.
	}
	defer db.Close()
	logger.Info("Database connection successful!")
	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	os.LookupEnv("PATH")
	mux := app.mount()
	logger.Fatal(app.run(mux))
}
