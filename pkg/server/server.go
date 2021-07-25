package server

import (
	"fmt"
	"log"
	"net/http"
)

// serverBuilder is a struct helper to build a server
type serverBuilder struct {
	port int
	// handlers is a map of http.HandlerFunc indexed by the path they handle
	handlers map[string]http.HandlerFunc
}

func defaultHandlerFunc() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Golang server running..."))
		if err != nil {
			log.Printf("failed to write response to %v", request)
		}
	}
}

// NewServerBuilder build a serverBuilder with default values for port (8080) and defaultHandlers (handler for "/").
// The handler can be overriden by the serverBuilder.Handle method.
func NewServerBuilder() *serverBuilder {
	// Define default values for the serverBuilder
	var (
		defaultPort     = 8080
		defaultHandlers = map[string]http.HandlerFunc{"/": defaultHandlerFunc()}
	)

	// Returns a serverBuilder with default values
	return &serverBuilder{
		port:     defaultPort,
		handlers: defaultHandlers,
	}
}

// Port sets the port the server listens on
func (s *serverBuilder) Port(port int) *serverBuilder {
	s.port = port
	return s
}

// Handle defines a function to handle a given path.
// Since serverBuilder.handlers is a map, one can only have one function for a given path.
// Subsequent calls to a given path will override the function for that path.
func (s *serverBuilder) Handle(path string, f http.HandlerFunc) *serverBuilder {
	s.handlers[path] = f
	return s
}

// Build returns a http.Server with the current config from serverBuilder
func (s *serverBuilder) Build() http.Server {
	mux := http.NewServeMux()
	for path, handlerFunc := range s.handlers {
		mux.HandleFunc(path, handlerFunc)
	}

	serverAddr := fmt.Sprintf(":%d", s.port)

	return http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
}
