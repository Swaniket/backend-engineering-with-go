package store

import (
	"context"
	"database/sql"
	"errors"
)

// Declaring error to use through out the storage layer
var (
	ErrNotFound = errors.New("resource not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		UpdatePost(context.Context, int64, int64, *UpdatePostPayload) (*UpdatePostResponse, error)
		DeletePost(context.Context, int64) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetCommentsByPostId(context.Context, int64) ([]Comment, error)
	}
}

// Constructor function for Storage
func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentsStore{db},
	}
}
