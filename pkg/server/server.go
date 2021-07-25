package server

import (
	"fmt"
	"net/http"
)

type serverBuilder struct {
	port     int
	handlers []handler
}

type handler struct {
	path string
	function http.HandlerFunc
}

func NewServerBuilder() *serverBuilder {
	return &serverBuilder{
		port:     8080,
		handlers: nil,
	}
}

func (s *serverBuilder) Port(port int) *serverBuilder {
	s.port = port
	return s
}

func (s *serverBuilder) Handle(path string, f http.HandlerFunc) *serverBuilder {
	s.handlers = append(s.handlers, handler{
		path:     path,
		function: f,
	})
	return s
}

func (s *serverBuilder) Build() http.Server {
	mux := http.NewServeMux()
	for _, handler := range s.handlers {
		mux.HandleFunc(handler.path, handler.function)
	}

	serverAddr := fmt.Sprintf(":%d", s.port)

	return http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
}
