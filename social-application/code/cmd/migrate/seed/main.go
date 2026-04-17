package main

import (
	"log"

	"github.com/Swaniket/social/internal/db"
	"github.com/Swaniket/social/internal/env"
	"github.com/Swaniket/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewPostgresStorage(conn)
	db.Seed(store, conn)
}
