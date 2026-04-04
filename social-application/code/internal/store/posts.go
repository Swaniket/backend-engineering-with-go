package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}

type UpdatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostResponse struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan( // Scan will get the returning values and assign them to the post struct
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetById(ctx context.Context, postID int64) (*Post, error) {
	query := `
		SELECT id, user_id, title, content, tags, created_at, updated_at 
		FROM posts 
		WHERE id = $1
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, postID).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) UpdatePost(ctx context.Context, postID int64, updatePost *UpdatePostPayload) (*UpdatePostResponse, error) {
	query := `
		UPDATE posts
		SET title = $1, content = $2
		WHERE id=$3
		RETURNING id, title, content, created_at, updated_at
	`

	var post UpdatePostResponse
	err := s.db.QueryRowContext(ctx, query, updatePost.Title, updatePost.Content, postID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostStore) DeletePost(ctx context.Context, postID int64) error {
	query := `
		DELETE FROM posts
		WHERE id=$1
	`

	// We are using ExecContext because we don't want to return anything
	res, err := s.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
