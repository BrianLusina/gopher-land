package gohttp

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	// concatenates hello with the body
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}
