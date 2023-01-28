package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// serverBuilder is the interface for concrete builders to implement
type serverBuilder interface {
	SetPort(port int) serverBuilder
	WithHandlerFunc(path string, f http.HandlerFunc) serverBuilder
	WithPathHandler(pathHandler PathHandler) serverBuilder
	Build() http.Server
}

// defaultServerBuilder is a concrete serverBuilder that implements the latter
type defaultServerBuilder struct {
	port int
	// handlers is a map of http.HandlerFunc indexed by the path they handle
	handlers map[string]http.HandlerFunc
}

// Define default values for the defaultServerBuilder
const (
	defaultPort     int = 8080
	defaultTemplate     = `<h1>Golang server is running...</h1>
<p>This is the default home page if you don't configure your server. Still, awesome that you're here, don't you agree?</p>
<p>You're in <code>{{.Path}}</code></p>`
)

var defaultHandlers = map[string]http.HandlerFunc{"/": defaultHandler}

// NewDefaultServerBuilder exposes the defaultServerBuilder with default values for port (8080) and defaultHandlers (handler for
// "/"). The handler can be overridden by the serverBuilder.WithHandlerFunc method.
func NewDefaultServerBuilder() *defaultServerBuilder {
	// Returns a defaultServerBuilder with default values
	return &defaultServerBuilder{
		port:     defaultPort,
		handlers: defaultHandlers,
	}
}

// SetPort sets the port the server listens on
func (b *defaultServerBuilder) SetPort(port int) serverBuilder {
	b.port = port
	return b
}

// WithHandlerFunc defines a function to handle a given path.
// Since defaultServerBuilder.handlers is a map, one can only have one function for a given path.
// Subsequent calls to a given path will override the function for that path.
func (b *defaultServerBuilder) WithHandlerFunc(path string, f http.HandlerFunc) serverBuilder {
	b.handlers[path] = f
	return b
}

// WithPathHandler defines several handlers for the same base path using a PathHandler.
func (b *defaultServerBuilder) WithPathHandler(pathHandler PathHandler) serverBuilder {
	for path, handlerFunction := range pathHandler.handlers {
		fullPath := pathHandler.basePath + path
		b.handlers[fullPath] = handlerFunction
	}
	return b
}

// Build returns a http.Server with the current config from defaultServerBuilder
func (b *defaultServerBuilder) Build() http.Server {
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
	t, err := template.New("default").Parse(defaultTemplate)
	if err != nil {
		log.Printf("Error parsing default template: %v", err)
	}

	data := struct {
		Path string
	}{
		Path: request.URL.Path,
	}

	t.Execute(writer, data)
}
