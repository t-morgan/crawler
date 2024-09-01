package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("no website provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		fmt.Printf("unable to parse maxConcurrency %s: %v\n", args[1], err)
		os.Exit(1)
	}

	maxPages, err := strconv.ParseInt(args[2], 10, 0)
	if err != nil {
		fmt.Printf("unable to parse maxPages %s: %v\n", args[2], err)
		os.Exit(1)
	}

	websiteURL := args[0]
	fmt.Printf(
		"starting crawl of: %v\nmaxConcurrency: %d\nmaxPages: %d\n",
		websiteURL,
		maxConcurrency,
		maxPages,
	)

	baseURL, err := url.Parse(websiteURL)
	if err != nil {
		fmt.Printf("unable to parse website URL: %v\n", err)
		os.Exit(1)
	}

	pages := map[string]int{}
	cfg := config{
		pages:              pages,
		maxPages:           int(maxPages), // because ParseInt thinks this is int64
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(websiteURL)
	cfg.wg.Wait()

	printReport(pages, websiteURL)
}
