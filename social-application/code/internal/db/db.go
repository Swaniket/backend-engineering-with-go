package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	db.SetMaxIdleConns(maxIdleConns)

	// Creating a context timeout for the DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Inline error handling for pinging the database
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
