package main

type Counter struct {
	total int
}

func(c *Counter)Inc() {
	c.total++
}

func(c Counter)Value() int {
	return c.total
}