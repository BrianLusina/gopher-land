package gohttp

import (
	"io"
	"log"
	"net/http"
)

type handler struct {
	client http.Client
	url    string
}

// getBody demonstrates a typical client handler to make a request and handle the body of the response. The body should be
// closed otherwise a memory leak is experiences. Since the Body is an io.ReadCloser, i.e. implements both the io.Reader and io.Closer, this body must be closed to prevent a leak. if this is not closed, the application will keep this in memory even though it's not utilized and the GC will fail to clean it up, causing other clients to not be able to re-use the TCP connection in the worst case
func (h handler) getBody() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to closes response: %v\n", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// getStatusCode demonstrates the same as getBody, but this is only interested in the status code of a request made by the client
// even in this scenario, the body still has to be closed to prevent a memory leak.
func (h handler) getStatusCode(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to closes response: %v\n", err)
		}
	}()

	// Read the response body even if it's not going to be used. Why? This is in the case that this method is called repeatedly which causes the HTTP transport to be closed preventing the TCP connection from being re-used. This follows from the face that if the body is closed without a read, the default HTTP transport closes the connection. If it is closed
	// following a read, then the transport won't close the connection and it will be re-used.
	// This code is more efficient than using io.ReadAll as it reads the body but discards it without any copy.
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.StatusCode, nil
}
