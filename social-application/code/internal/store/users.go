package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create(ctx context.Context) error {
	return nil
}
