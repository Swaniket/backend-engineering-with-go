package main

import (
	"net/http"
)

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("internal server error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("bad request: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) NotFoundError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("not found error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound, "Not found")
}

func (app *application) ConflictError(w http.ResponseWriter, r *http.Request, err error) {
	// log.Printf("conflict error: %s | path: %s | error: %s", r.Method, r.URL.Path, err.Error())
	app.logger.Errorw("conflict error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusConflict, "Resource Conflict")
}
