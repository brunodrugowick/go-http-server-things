package server

import (
	"fmt"
	"html/template"
	"net/http"
)

// builder is the interface for concrete builders to implement
type builder interface {
	SetPort(port int) builder
	WithHandler(path string, f http.HandlerFunc) builder
	Build() http.Server
}

// defaultBuilder is a concrete builder that implements the latter
type defaultBuilder struct {
	port int
	// handlers is a map of http.HandlerFunc indexed by the path they handle
	handlers map[string]http.HandlerFunc
}

// NewDefaultBuilder exposes the defaultBuilder with default values for port (8080) and defaultHandlers (handler for
// "/"). The handler can be overridden by the builder.WithHandler method.
func NewDefaultBuilder() *defaultBuilder {
	// Define default values for the defaultBuilder
	var (
		defaultPort     = 8080
		defaultHandlers = map[string]http.HandlerFunc{"/": defaultHandler}
	)

	// Returns a defaultBuilder with default values
	return &defaultBuilder{
		port:     defaultPort,
		handlers: defaultHandlers,
	}
}

// SetPort sets the port the server listens on
func (b *defaultBuilder) SetPort(port int) builder {
	b.port = port
	return b
}

// WithHandler defines a function to handle a given path.
// Since defaultBuilder.handlers is a map, one can only have one function for a given path.
// Subsequent calls to a given path will override the function for that path.
func (b *defaultBuilder) WithHandler(path string, f http.HandlerFunc) builder {
	b.handlers[path] = f
	return b
}

// Build returns a http.Server with the current config from defaultBuilder
func (b *defaultBuilder) Build() http.Server {
	mux := http.NewServeMux()
	for path, handlerFunc := range b.handlers {
		mux.HandleFunc(path, handlerFunc)
	}

	serverAddr := fmt.Sprintf(":%d", b.port)

	return http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
}

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("templates/default-home-page.html")

	data := struct {
		Path string
	}{
		Path: request.URL.Path,
	}

	t.Execute(writer, data)
}
