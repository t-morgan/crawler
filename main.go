package main

import (
	"fmt"
	"os"
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

	pages := map[string]int{}
	crawlPage(websiteURL, websiteURL, pages)

	for key, val := range pages {
		fmt.Printf("%s: %d\n", key, val)
	}
}
