package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	websiteURL := args[0]
	fmt.Printf("starting crawl of: %v\n", websiteURL)

	baseURL, err := url.Parse(websiteURL)
	if err != nil {
		fmt.Printf("unable to parse website URL: %v\n", err)
		os.Exit(1)
	}

	pages := map[string]int{}
	maxConcurrency := 10
	cfg := config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(websiteURL)
	cfg.wg.Wait()

	for key, val := range pages {
		fmt.Printf("%s: %d\n", key, val)
	}
}
