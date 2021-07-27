package handlers

import (
	"fmt"
	"net/http"
)

func HandlerRoot (w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, mainHTMLPage())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// For demonstration purposes only
func mainHTMLPage() string {
	return `<html><h1>Welcome</h1><p>Try to hit <a href="/hello?name=John">/hello?name=John</a> to see another handler in action.</html>`
}
