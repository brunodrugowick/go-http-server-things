package server

import (
	"net/http"
)

// pathHandlerBuilder is the interface for concrete builders to implement
type pathHandlerBuilder interface {
	WithHandlerFunc(subtree string, f http.HandlerFunc) pathHandlerBuilder
	Build() PathHandler
}

// defaultPathHandlerBuilder is a concrete pathHandlerBuilder that implements the latter
type defaultPathHandlerBuilder struct {
	basePath string
	handlers map[string]http.HandlerFunc
}

// NewDefaultPathHandlerBuilder constructs the defaultPathHandlerBuilder with a provided basePath
func NewDefaultPathHandlerBuilder(basePath string) *defaultPathHandlerBuilder {
	return &defaultPathHandlerBuilder{
		basePath: basePath,
		handlers: map[string]http.HandlerFunc{},
	}
}

// PathHandler is the struct to be built
type PathHandler struct {
	basePath	string
	handlers	map[string]http.HandlerFunc
}

// WithHandlerFunc specifies a function that handles a subtree of the basePath of PathHandler
func (hb *defaultPathHandlerBuilder) WithHandlerFunc(subtree string, f http.HandlerFunc) pathHandlerBuilder {
	hb.handlers[subtree] = f
	return hb
}

// Build returns a configured PathHandler
func (hb *defaultPathHandlerBuilder) Build() PathHandler {
	return PathHandler{
		basePath: hb.basePath,
		handlers: hb.handlers,
	}
}