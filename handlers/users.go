package handlers

import (
	"encoding/json"
	"github.com/brunodrugowick/go-http-server-things/pkg/server"
	"net/http"
	"strings"
)

func UsersPathHandler() server.PathHandler {
	// It's imperative to understand how Mux handles the patterns given to it.
	// See https://pkg.go.dev/net/http#ServeMux
	return server.NewDefaultPathHandlerBuilder("/users").
		WithHandlerFunc("", handleUsersCollection).
		WithHandlerFunc("/", handleUserResource).
		Build()
}

func handleUsersCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		serveUsersCollection(w)
	case http.MethodPost:
		createNewUser()
	}
}

func handleUserResource(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")
	// TODO Validations on userID?

	switch r.Method {
	case http.MethodGet:
		serveUserResource(w, userID)
	case http.MethodPost:
		editUserResource(w, userID)
	}
}

func serveUsersCollection(w http.ResponseWriter) {
	response := []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		{
			Name:  "John",
			Email: "john@email.com",
		},
		{
			Name:  "Mark",
			Email: "mark@email.com",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func createNewUser() {
	// Not implemented
}

func serveUserResource(w http.ResponseWriter, userID string) {
	response := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  userID,
		Email: userID + "@email.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func editUserResource(w http.ResponseWriter, userID string) {
	// Not implemented
}
