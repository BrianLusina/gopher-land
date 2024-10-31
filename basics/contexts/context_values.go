package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type key string

type message struct{}

const isValidHostKey = "isValidHost"

func checkValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validHost := r.Host == "acme"
		ctx := context.WithValue(r.Context(), isValidHostKey, validHost)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(ctx context.Context, ch chan message) error {
	for {
		select {
		case msg := <-ch:
			// do something with message
			_ = fmt.Sprintf("Message: %v", msg)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
