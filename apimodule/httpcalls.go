package orcidapi

import (
	"bytes"
	"net/http"
)

// MakePostRequest sends an HTTP POST request to the specified URL with the given data
// and returns the response and any error encountered.
func makePostRequest(url string, contentType string, body []byte) (*http.Response, error) {
	// Create a new POST request with the provided body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", contentType)

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// makes a get request
func makeGetRequest(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
	if err != nil {
		return nil, err
	}
	req.Header = headers
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
