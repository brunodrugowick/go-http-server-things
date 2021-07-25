package main

import (
	"github.com/brunodrugowick/go-http-server-things/pkg/server"
	"log"
	"net/http"
)

func main() {

	server := server.NewServerBuilder().
		Port(8085).
		Handle("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("You've reached the server"))
	}).
		Build()

	log.Fatal(server.ListenAndServe())
}
