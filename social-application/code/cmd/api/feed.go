package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// @TODO: Pagination, Filters
	ctx := r.Context()
	userId := int64(35)

	feed, err := app.store.Posts.GetUserFeed(ctx, userId)
	if err != nil {
		app.InternalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.InternalServerError(w, r, err)
	}
}
