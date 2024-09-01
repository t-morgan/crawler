package main

import (
	"fmt"
	"os"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) map[string]int {
	if !strings.HasPrefix(rawCurrentURL, rawBaseURL) {
		return pages
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL]++
		return pages
	}

	pages[normalizedURL] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(html)

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

	return pages
}
