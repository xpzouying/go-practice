package main

import (
	"fmt"
	"sync"
)

func append1() {
	buf := make([]int, 0, 10)
	var mu sync.Mutex

	for i := 0; i < 20; i++ {
		// append
		mu.Lock()
		buf = append(buf, i)
		mu.Unlock()
		fmt.Printf("append buf: %v\n", buf)

		// remove
		mu.Lock()
		buf = append(buf[:0], buf[1:]...)
		mu.Unlock()
		fmt.Printf("remove buf: %v, buf addr: %p, len(buf)=%v, cap(buf)=%v\n",
			buf, &buf, len(buf), cap(buf))
	}
}

func append2() {
	buf := make([]int, 0, 10)
	var mu sync.Mutex

	for i := 0; i < 20; i++ {
		// append
		mu.Lock()
		buf = append(buf, i)
		mu.Unlock()
		fmt.Printf("append buf: %v\n", buf)

		// remove
		mu.Lock()
		// buf = append(buf[:0], buf[1:]...)
		// buf = append(buf[:0], buf[1:]...)
		buf = buf[1:]
		mu.Unlock()
		fmt.Printf("remove buf: %v, buf addr: %p, len(buf)=%v, cap(buf)=%v\n",
			buf, &buf, len(buf), cap(buf))
	}
}

func main() {
	fmt.Println("append1:")
	append1()

	fmt.Println("append2:")
	append2()
}
