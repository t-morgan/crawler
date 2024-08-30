package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme wss",
			inputURL: "wss://ws.boot.dev/socket",
			expected: "ws.boot.dev/socket",
		},
		{
			name:     "url without scheme",
			inputURL: "i.have.no.scheme/to/hold/me/down",
			expected: "i.have.no.scheme/to/hold/me/down",
		},
		{
			name:     "remove query parameters",
			inputURL: "https://api.boot.dev/path?sort=quality",
			expected: "api.boot.dev/path",
		},
		{
			name:     "remove query parameters",
			inputURL: "https://api.boot.dev/path?sort=quality",
			expected: "api.boot.dev/path",
		},
		{
			name:     "remove fragment",
			inputURL: "https://blog.boot.dev/path#fragment",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "maintain port",
			inputURL: "https://blog.boot.dev:3000/path",
			expected: "blog.boot.dev:3000/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
