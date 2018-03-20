package main

import (
	"sync"
	"testing"
)

func BenchmarkMutexSlice(b *testing.B) {
	buf := make([]int, 10)
	var mu sync.Mutex

	for i := 0; i < b.N; i++ {
		// append
		mu.Lock()
		buf = append(buf, i)
		mu.Unlock()

		// remove
		mu.Lock()
		buf = append(buf[:0], buf[1:]...)
		mu.Unlock()
	}
}

func BenchmarkChannel(b *testing.B) {
	c := make(chan int, 10)

	for i := 0; i < b.N; i++ {
		c <- i
		<-c
	}
}
