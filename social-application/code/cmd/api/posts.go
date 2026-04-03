package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Swaniket/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=3000"`
	Tags    []string `json:"tags"`
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

	if err := WriteJson(w, http.StatusCreated, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "postID")

	// Parse the ID param & convert it into int64
	postID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	// Get Post by ID from store
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

	if err := WriteJson(w, http.StatusOK, post); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
