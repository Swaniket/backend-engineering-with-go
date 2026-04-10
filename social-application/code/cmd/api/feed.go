package main

import (
	"net/http"

	"github.com/Swaniket/social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// @TODO: Pagination, Filters
	// Default Pagination Values
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	// Parse the request to get the Pagination params
	fq, err := fq.ParsePaginationRequest(r)
	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	userId := int64(32)

	feed, err := app.store.Posts.GetUserFeed(ctx, userId, fq)
	if err != nil {
		app.InternalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.InternalServerError(w, r, err)
	}
}
