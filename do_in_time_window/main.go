package main

import (
	"fmt"
	"time"
)

func generateData(c chan<- string) {
	defer func() {
		close(c)
		c = nil
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)

		c <- fmt.Sprintf("data%d", i)
	}
}

// pack data into data package, data package is simple []string
func packData(srcc <-chan string, destc chan<- []string) {
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
	defer close(destc)

	var ss []string
	for {
		select {
		case <-t.C:
			destc <- ss
			fmt.Println("send data:", ss)

			ss = []string{}
		case s, ok := <-srcc:
			if !ok {

				// send buffered data
				if len(ss) > 0 {
					destc <- ss
					fmt.Println("send buffered data:", ss)
				}

				fmt.Println("no src data anymore")
				return
			}
			ss = append(ss, s)
		}
	}
}

func main() {
	datac := make(chan string, 100)
	destc := make(chan []string, 100)

	go generateData(datac)
	go packData(datac, destc)

	for rec := range destc {
		fmt.Println("received data:", rec)
	}
}
