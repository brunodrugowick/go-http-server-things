package handlers

import (
	"encoding/json"
	"net/http"
)

func HandlerHelloWithQueryParam(w http.ResponseWriter, r *http.Request) {
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
