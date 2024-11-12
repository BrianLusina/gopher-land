package gohttp

import (
	"net"
	"net/http"
	"time"
)

func setupClient() *http.Client {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second,
		},
	}
	return client
}

func setupServer(handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		IdleTimeout:       time.Second,
		Handler:           http.TimeoutHandler(handler, time.Second, "foo"),
	}

	return server
}
