package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Swaniket/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type postKey string

const postCtx postKey = "postFromContext"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=3000"`
	Tags    []string `json:"tags"`
}

type UpdatePostClientPayload struct {
	Title   string `json:"title" validate:"omitempty,max=100"`
	Content string `json:"content" validate:"omitempty,max=3000"`
}

func (app *application) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the post payload from the request
	var payload CreatePostPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		// Todo: Change after auth
		UserID: 1,
		Tags:   payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	comments, err := app.store.Comments.GetCommentsByPostId(r.Context(), post.ID)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	post.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) UpdatePostByIDHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	// Get the payload from request body
	var payload UpdatePostClientPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	updatePost := &store.UpdatePostPayload{
		Title: func() string {
			if payload.Title == "" {
				return post.Title
			}
			return payload.Title
		}(),
		Content: func() string {
			if payload.Content == "" {
				return post.Content
			}
			return payload.Content
		}(),
	}

	updatedPost, err := app.store.Posts.UpdatePost(r.Context(), post.ID, post.Version, updatePost)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.ConflictError(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, updatedPost); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "postID")
	// Parse the ID param & convert it into int64
	postID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.store.Posts.DeletePost(ctx, postID); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.NotFoundError(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent) // Return nothing after successful deletion
}

// This is fetch the post and put it inside of the context
func (app *application) postsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idParam := chi.URLParam(r, "postID")
		// Parse the ID param & convert it into int64
		postID, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.InternalServerError(w, r, err)
			return
		}

		// Get the post from the postID
		post, err := app.store.Posts.GetById(ctx, postID)
		if err != nil {
			// Handle not found error as well
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.NotFoundError(w, r, err)
			default:
				app.InternalServerError(w, r, err)
			}
			return
		}

		// Create a new context and put the post inside of that context
		ctx = context.WithValue(ctx, postCtx, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromCtx(r *http.Request) *store.Post {
	post, _ := r.Context().Value(postCtx).(*store.Post)
	return post
}
