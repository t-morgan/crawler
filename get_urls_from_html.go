package main

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse html: %w", err)
	}

	urls := []string{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					val := a.Val

					valURL, err := url.Parse(a.Val)
					if err != nil {
						fmt.Println(err)
						break
					}

					ext := filepath.Ext(valURL.Path)
					if ext != "" && ext != ".html" && ext != ".htm" {
						break
					}

					if strings.HasPrefix(val, "/") {
						urls = append(urls, rawBaseURL+val)
					} else if strings.HasPrefix(val, "..") {
						splitBaseURL := strings.Split(rawBaseURL, "/")
						splitVal := strings.Split(val, "/")
						count := 0
						newVal := ""
						for _, segment := range splitVal {
							if segment == ".." {
								count++
							} else {
								newVal += "/" + segment
							}
						}
						relativeBaseURL := strings.Join(splitBaseURL[:len(splitBaseURL)-count], "/")
						urls = append(urls, relativeBaseURL+newVal)
					} else {
						urls = append(urls, val)
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls, nil
}
