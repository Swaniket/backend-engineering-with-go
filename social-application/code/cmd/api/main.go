package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mountRoutes()

	log.Fatal(app.run(mux))
}
