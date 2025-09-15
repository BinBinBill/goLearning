package main

import (
	"sync"
)

type safeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *safeCounter) inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func (c *safeCounter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
