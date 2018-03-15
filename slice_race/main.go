package main

import (
	"fmt"
	"sync"
)

func appendToSlice(s []int, b int) {
	for i := b; i < b+100; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	// s := make([]int, 0, 100) // -race complains because the first appends will modify the same backing array
	// s := make([]int, 0) // -race doesn't complain because the first append will make a new backing array
	var s []int // -race doesn't complain because the first append will make a new backing array

	wg.Add(3)
	go func(s []int) {
		appendToSlice(s, 1)
		wg.Done()
	}(s)
	go func(s []int) {
		appendToSlice(s, 100)
		wg.Done()
	}(s)
	go func(s []int) {
		appendToSlice(s, 200)
		wg.Done()
	}(s)

	wg.Wait()

}
