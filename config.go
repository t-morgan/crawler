package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) isMaxPagesReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}
