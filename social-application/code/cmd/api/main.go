package main

import (
	"log"
	"os"

	"github.com/Swaniket/social/internal/db"
	"github.com/Swaniket/social/internal/env"
	"github.com/Swaniket/social/internal/store"
)

func main() {
	// Having "postgres://user:adminpassword@localhost/social?sslmode=disable" - as the default
	// We are going to run postgres with docker for local, and for that we don't need DB_ADDR env variable
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err) // Don't want to bootup the server if the db connection fails.
	}

	defer db.Close()
	log.Printf("Database connection successful!")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
