package main

import "net/url"

func normalizeURL(inputURL string) (string, error) {
	u, err := url.Parse(inputURL)

	if err != nil {
		return "", err
	}

	cleanedUrl := u.Host + u.Path

	return cleanedUrl, nil
}
