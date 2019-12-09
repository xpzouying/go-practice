package main

import (
	"fmt"
	"sync"
	"time"
)

func deferInLoop1() {
	fmt.Println("defer in loop 1: ")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("defer func\n")
		fmt.Printf("number: %d\n", i)
	}

	fmt.Println("finish")
}

func deferInLoop2() {
	fmt.Println("defer in loop 2: ")
	for i := 0; i < 10; i++ {
		func() {
			defer fmt.Printf("defer func\n")
			fmt.Printf("number: %d\n", i)
		}()
	}

	fmt.Println("finish")
}

func deferInLoop3() {
	fmt.Println("defer in loop 3: ")

	var wg sync.WaitGroup
	wg.Add(10)

	// FIXME: defer wg.Done() is error
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("defer func\n")
			wg.Done()
		}()
		fmt.Printf("number: %d\n", i)
	}

	wg.Wait()

	fmt.Println("finish")
}

func deferCallTime1() {
	begin := time.Now()
	fmt.Println("now: ", begin)

	defer func(t time.Time) {
		fmt.Println("defer time: ", t)
	}(time.Now())

	time.Sleep(time.Second)

	fmt.Println("end: ", time.Now())
}

func deferCallTime2() {
	begin := time.Now()
	fmt.Println("now: ", begin)

	defer func() {
		fmt.Println("defer time: ", time.Now())
	}()

	time.Sleep(time.Second)

	fmt.Println("end: ", time.Now())
}

func main() {
	deferInLoop1()
	deferInLoop2()

	// FIXME: deferInLoop3() have bug
	// deferInLoop3()

	deferCallTime1()
	deferCallTime2()
}
