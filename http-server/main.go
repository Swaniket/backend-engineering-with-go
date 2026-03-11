package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// Basic routing
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("index page"))
// 			return
// 		case "/users":
// 			w.Write([]byte("users page"))
// 			return
// 		default:
// 			w.Write([]byte("404 page"))
// 			return
// 		}

// 	default:
// 		w.Write([]byte("404 page not found"))
// 		return

// 	}
// }

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list..."))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User"))
}

func main() {
	api := &api{addr: ":8080"}

	// Init ServeMux - Which is a router library built in
	mux := http.NewServeMux() // It implements a handler interface

	server := &http.Server{
		Addr: api.addr,
		// Handler: api,
		Handler: mux,
		// We can also add other configs
	}

	// Adding routes using mux
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
