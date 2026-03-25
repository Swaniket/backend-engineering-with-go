package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context) error
	}
	Users interface {
		Create(context.Context) error
	}
}

// Constructor function for Storage
func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UserStore{db},
	}
}
