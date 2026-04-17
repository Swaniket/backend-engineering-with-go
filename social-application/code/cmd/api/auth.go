package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/Swaniket/social/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

func (app *application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := ReadJSON(w, r, &payload); err != nil {
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

	plainToken := uuid.New().String()

	// Encrypt the token and store it into the DB -> sent users the plain token
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	if err := app.store.Users.CreateAndInvite(ctx, user, hashToken, app.config.mail.exp); err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.BadRequestError(w, r, err)
			return
		case store.ErrDuplicateUsername:
			app.BadRequestError(w, r, err)
			return
		default:
			app.InternalServerError(w, r, err)
			return
		}
	}

	userWithToken := UserWithToken{
		User:  user,
		Token: plainToken,
	}

	//@TODO: Send the email - also rollback the user storing if it fails using SQL transactions

	// Send success after account is created
	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
