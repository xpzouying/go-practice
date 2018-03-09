package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
}

func (c *counter) Inc1() {
	c.count++
}

func (c *counter) Inc2() {
	c.count++
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
