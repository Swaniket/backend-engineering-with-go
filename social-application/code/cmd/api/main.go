package main

import (
	"log"
	"os"

	"github.com/Swaniket/social/internal/env"
	"github.com/Swaniket/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
