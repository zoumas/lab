package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       *sync.Mutex
	counters map[string]int
}

func NewContainer() *Container {
	return &Container{
		mu:       &sync.Mutex{},
		counters: make(map[string]int),
	}
}

func (c *Container) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := NewContainer()
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.Inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10_000)
	go doIncrement("a", 10_000)
	go doIncrement("b", 10_000)

	wg.Wait()
	fmt.Println(c.counters)
}
