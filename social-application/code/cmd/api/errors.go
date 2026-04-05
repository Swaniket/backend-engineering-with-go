package main

import (
	"log"
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Later we will have proper logging
	log.Printf("internal server error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	// Later we will have proper logging
	log.Printf("bad request: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) NotFoundError(w http.ResponseWriter, r *http.Request, err error) {
	// Later we will have proper logging
	log.Printf("not found error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusNotFound, "Not found")
}

func (app *application) ConflictError(w http.ResponseWriter, r *http.Request, err error) {
	// Later we will have proper logging
	log.Printf("conflict error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusConflict, "Resource Conflict")
}
