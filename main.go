package main

import (
	"encoding/json"
	"fmt"
	"github.com/brunodrugowick/go-http-server-things/pkg/server"
	"log"
	"net/http"
)

func main() {

	handlerRoot := func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, mainHTMLPage())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	handlerName := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		response := []struct {
			Say string `json:"Say"`
			To  string `json:"To"`
		}{
			{
				Say: "Hello",
				To:  name,
			},
			{
				Say: "Bye",
				To:  name,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	// build a basic server from the server package
	srv := server.NewDefaultBuilder().
		// to listen on 8085
		SetPort(8085).
		// with a basic handler for "/"
		WithHandler("/", handlerRoot).
		WithHandler("/hello", handlerName).
		Build()

	// starts the server
	log.Fatal(srv.ListenAndServe())
}

// For demonstration purposes only
func mainHTMLPage() string {
	return `<html>
	<h1>Welcome</h1>
	<p>Try to hit <a href="/hello?name=John">/hello?name=John</a> to see another handler in action.
	</html>`
}
