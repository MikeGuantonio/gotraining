package main

import "sync"

type Counter struct {
	mu sync.Mutex
	total int
}

func NewCounter() *Counter {
	return &Counter{}
}

func(c *Counter)Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.total++
}

func(c *Counter)Value() int {
	return c.total 
}