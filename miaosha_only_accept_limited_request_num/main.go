package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type request struct {
	uid     int
	reqTime time.Time
}

func main() {
	// done for finish
	done := make(chan struct{})

	// make channel to accept request, only accept top 3
	reqChan := make(chan request, 3)

	go func() {
		var wg sync.WaitGroup
		// send request into reqChan channel
		for i := 0; i < 10; i++ {
			req := request{uid: i}
			wg.Add(1)

			go func() {
				interval := rand.Intn(3000) + 1000
				time.Sleep(time.Duration(interval) * time.Millisecond)

				req.reqTime = time.Now()
				reqChan <- req
				wg.Done()
			}()
		}

		wg.Wait()
		close(reqChan)

		done <- struct{}{}
	}()

	// output
	for {
		select {
		case <-done:
			return
		default:
		case r, ok := <-reqChan:
			if !ok {
				reqChan = nil
				continue
			}
			fmt.Printf("received request: %v, %v\n", r.uid, r.reqTime.String())
		}
	}
}
