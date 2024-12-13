package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {

	res, err := http.Get(rawURL)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return "", fmt.Errorf("response failed with error code: %d", res.StatusCode)
	}
	if err != nil {
		return "", err
	}

	return string(body), nil
}
