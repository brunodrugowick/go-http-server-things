package main

import (
	"github.com/brunodrugowick/go-http-server-things/handlers"
	"github.com/brunodrugowick/go-http-server-things/pkg/server"
	"log"
)

func main() {

	// Creates a default server in port 9090... just because.
	anotherSrv := server.NewDefaultServerBuilder().SetPort(9090).Build()
	go func() {
		err := anotherSrv.ListenAndServe()
		if err != nil {
			log.Printf("Failed to start server in 9090 because %v", err)
		}
	}()

	// build a basic server from the server package
	srv := server.NewDefaultServerBuilder().
		// to listen on 8085
		SetPort(8085).
		// with a basic handler for "/"
		WithHandlerFunc("/", handlers.HandlerRoot).
		WithHandlerFunc("/hello", handlers.HandlerHelloWithQueryParam).
		WithPathHandler(handlers.UsersPathHandler()).
		Build()

	// starts the server
	log.Fatal(srv.ListenAndServe())
}
