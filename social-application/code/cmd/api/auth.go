package main

import (
	"net/http"

	"github.com/Swaniket/social/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := ReadJSON(w, r, payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	// Hash the user password
	if err := user.Password.Set(payload.Password); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	// Store the user
	ctx := r.Context()

	if err := app.store.Users.CreateAndInvite(ctx, user, "Token-123"); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	// Send the email - also rollback the user storing if it fails using SQL transactions

	// Send success after account is created
	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
