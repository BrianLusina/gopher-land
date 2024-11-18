package gohttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDurationClientGet(t *testing.T) {
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"duration": 314}`))
		}),
	)

	defer srv.Close()

	client := NewDurationClient()
	duration, err := client.GetDuration(srv.URL, coordinates{Lat1: 51.551261, Lng1: -0.1221146, Lat2: 51.57, Lng2: -.13})
	if err != nil {
		t.Fatal(err)
	}

	if duration != 314*time.Second {
		t.Errorf("expected 314 seconds, got %v", duration)
	}
}
