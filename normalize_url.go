package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse url: %w", err)
	}

	port := parsedURL.Port()
	if port != "" {
		port = ":" + port
	}
	normalized := parsedURL.Hostname() + port + parsedURL.Path

	return normalized, nil
}
