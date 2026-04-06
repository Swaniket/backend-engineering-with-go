package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Swaniket/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "userFromContext"

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	followedUser := getUserFromCtx(r)
	ctx := r.Context()

	// Revert back to auth userId from ctx
	var payload FollowUser
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := app.store.Followers.Follow(ctx, followedUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.ConflictError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

func (app *application) UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromCtx(r)
	ctx := r.Context()

	// Revert back to auth userId from ctx
	var payload FollowUser
	if err := ReadJSON(w, r, &payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := app.store.Followers.Unfollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.ConflictError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// Get the userId -> Fetch the user -> put it into the context
func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
