package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// We can either do it manually like this
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"status": "ok"}`))

	// Or
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	err := WriteJson(w, http.StatusOK, data)

	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
