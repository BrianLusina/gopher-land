package gohttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DurationClient struct {
	client *http.Client
}

type coordinates struct {
	Lat1, Lng1, Lat2, Lng2 float64
}

type response struct {
	Duration int64
}

func NewDurationClient() DurationClient {
	c := setupClient()
	return DurationClient{
		client: c,
	}
}

func (c DurationClient) GetDuration(url string, coordinates coordinates) (time.Duration, error) {
	resp, err := c.client.Post(url, "application/json", buildRequestBody(coordinates))
	if err != nil {
		return 0, err
	}

	return parseResponseBody(resp.Body)
}

func buildRequestBody(coordinates coordinates) *bytes.Buffer {
	reqBody, err := json.Marshal(coordinates)
	if err != nil {
		fmt.Printf("Error while marshalling: %+v\n", err)
		return nil
	}
	return bytes.NewBuffer(reqBody)
}

func parseResponseBody(responseBody io.ReadCloser) (time.Duration, error) {
	body := []byte{}
	_, err := responseBody.Read(body)
	if err != nil {
		return 0, err
	}

	var resp *response
	if err := json.Unmarshal(body, resp); err != nil {
		return 0, err
	}

	return time.Duration(resp.Duration), nil

}
