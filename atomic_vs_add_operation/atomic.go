package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count int64
}

func (c *counter) Inc1() {
	atomic.AddInt64(&c.count, 1)
}

func (c *counter) Inc2() {
	atomic.AddInt64(&c.count, 1)
}

func main() {

	var c counter

	var wg sync.WaitGroup
	wg.Add(2000)
	go func() {
		for i := 0; i < 1000; i++ {
			c.Inc1()
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			c.Inc2()
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Printf("counter: %+v\n", c)
}
