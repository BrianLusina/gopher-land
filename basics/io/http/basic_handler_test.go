package gohttp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	testReq := httptest.NewRequest(http.MethodGet, "http://localhost", strings.NewReader("foo"))
	responseRecorder := httptest.NewRecorder()
	Handler(responseRecorder, testReq)

	if got := responseRecorder.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, go %s", got)
	}

	body, _ := io.ReadAll(wordy)
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	if http.StatusOK != responseRecorder.Result().StatusCode {
		t.FailNow()
	}
}
