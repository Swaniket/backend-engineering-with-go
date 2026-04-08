package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Declaring error to use through out the storage layer
var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 10 // 10 Second query timeout duration
)

type Storage struct {
	DB    *sql.DB
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		UpdatePost(context.Context, int64, int64, *UpdatePostPayload) (*UpdatePostResponse, error)
		DeletePost(context.Context, int64) error
		GetUserFeed(context.Context, int64) ([]PostForFeed, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetCommentsByPostId(context.Context, int64) ([]Comment, error)
	}
	Followers interface {
		Follow(context.Context, int64, int64) error
		Unfollow(context.Context, int64, int64) error
	}
}

// Constructor function for Storage
func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		DB:        db,
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentsStore{db},
		Followers: &FollowerStore{db},
	}
}
