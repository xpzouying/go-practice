package main

import (
	"fmt"
	"sync"
)

func addWithoutMutex() {
	var m sync.Map

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			m.Store(i, i)
		}(i)
	}

	wg.Wait()

	fn := func(key, value interface{}) bool {
		fmt.Printf("key=%d, value=%d\n", key, value)
		return true
	}

	m.Range(fn)
}

func main() {
	addWithoutMutex()
}
