package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/time/rate"
)

type Task struct {
	name   string
	result chan result
}

type result struct {
	finish bool
	error  string
}

func producer(producerName string, taskQueue chan<- *Task) {
	for range time.Tick(100 * time.Millisecond) {
		taskID := rand.Int31n(100)
		taskName := fmt.Sprintf("%s-taskid-%d", producerName, taskID)

		task := &Task{
			name:   taskName,
			result: make(chan result),
		}
		taskQueue <- task

		result := <-task.result
		if result.finish {
			log.Printf("finish task: %s", task.name)
		} else {
			// log.Printf("task failed: %s, error: %s", task.name, result.error)
		}
	}
}

func consumer(taskQueue <-chan *Task) {
	every := rate.Every(100 * time.Millisecond)
	lmt := rate.NewLimiter(every, 10)

	for task := range taskQueue {
		// log.Printf("consumer got task: %s", task.name)
		if ok := lmt.Allow(); !ok {
			task.result <- result{finish: false, error: "out of limit"}
			continue
		}

		task.result <- result{finish: true}
	}
}

func main() {
	// task queue
	taskQueue := make(chan *Task, 100)

	// start multiple producer
	for i := 0; i < 2; i++ {
		producerName := fmt.Sprintf("producer-%d", i)

		go producer(producerName, taskQueue)
	}

	// start single consumer worker
	consumer(taskQueue)
}
