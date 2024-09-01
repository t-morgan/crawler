package main

import (
	"fmt"
	"sort"
	"strings"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	reportPages := []page{}
	for url, count := range pages {
		reportPages = append(reportPages, page{URL: url, Count: count})
	}
	sort.Slice(reportPages, func(i, j int) bool {
		return reportPages[i].Count > reportPages[j].Count ||
			strings.Compare(reportPages[i].URL, reportPages[j].URL) < 0
	})

	for _, page := range reportPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}

type page struct {
	URL   string
	Count int
}
