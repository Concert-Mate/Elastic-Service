package helpers

import (
	"fmt"
	"io"
	"net/http"
)

func GetContentFromURL(url string) ([]byte, error) {
	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			err = cErr
		}
	}()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
