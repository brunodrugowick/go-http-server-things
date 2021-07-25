package main

import (
	"github.com/brunodrugowick/go-http-server-things/pkg/server"
	"log"
	"net/http"
)

func main() {

	// build a basic server from the server package
	srv := server.NewServerBuilder().
		// to listen on 8085
		Port(8085).
		// with a basic handler for "/"
		Handle("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("You've reached the server"))
		}).
		Build()

	// starts the server
	log.Fatal(srv.ListenAndServe())
}
