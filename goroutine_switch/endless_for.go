package main

import (
	"fmt"
	"runtime"
	"sync"
)

func runForever(id int) {
	fmt.Printf("id: %d\n", id)
	for {
	}
}
func main() {
	fmt.Println("MAXPROCS: ", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go runForever(i)
	}
	wg.Wait()
}
