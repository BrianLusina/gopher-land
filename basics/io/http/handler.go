package gohttp

import (
	"log"
	"net/http"
)

func demo_handler(w http.ResponseWriter, req *http.Request) {
	err := foo(req)
	if err != nil {
		http.Error(w, "foo", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte("all good"))
	w.WriteHeader(http.StatusCreated)
}

func foo(req *http.Request) error {
	log.Printf("Handling request :%v\n", req)
	return nil
}
