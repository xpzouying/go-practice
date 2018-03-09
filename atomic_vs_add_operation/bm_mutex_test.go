package main

import (
	"sync"
	"testing"
)

func BenchmarkMutexCounter(b *testing.B) {

	c := counter{}

	var wg sync.WaitGroup
	NUM := b.N
	wg.Add(NUM * 2)

	for i := 0; i < NUM; i++ {
		go func() {
			c.Inc1()
			wg.Done()
		}()

		go func() {
			c.Inc2()
			wg.Done()
		}()
	}

	wg.Wait()

	if c.count != NUM*2 {
		b.Errorf("number of counter not equal")
	}
}
