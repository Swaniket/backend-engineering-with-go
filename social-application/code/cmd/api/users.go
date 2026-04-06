package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Swaniket/social/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetByID(ctx, userID)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.NotFoundError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
