package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("unable to get url `%s`: %w", rawURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("error returned from get: %s", res.Status)
	}

	contentType := res.Header.Get("content-type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("unable to handle content type: %s", contentType)
	}

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response body: %w", err)
	}

	return string(html), nil
}
