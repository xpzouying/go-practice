package main

import (
	"fmt"
	"sync"
)

func addWithMutex() {
	m := make(map[int]int)
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	for k, v := range m {
		fmt.Printf("key=%d, value=%d\n", k, v)
	}
}

func main() {
	addWithMutex()
}
