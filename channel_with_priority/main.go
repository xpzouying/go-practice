package main

import (
	"fmt"
	"time"
)

// Msg is just string type
type Msg = string

var (
	msgQueue map[string]chan Msg
)

func prioritys() []string {
	return []string{"HIGH", "LOW"}
}

func main() {
	// init msg queue
	msgQueue := make(map[string]chan Msg, len(prioritys()))
	for _, p := range prioritys() {
		msgQueue[p] = make(chan Msg, 10)
	}

	// product msg
	go func() {
		msgQueue["HIGH"] <- "high msg1"
		msgQueue["LOW"] <- "low msg1"
		msgQueue["HIGH"] <- "high msg2"
		msgQueue["HIGH"] <- "high msg3"
		msgQueue["LOW"] <- "low msg2"

		for _, p := range prioritys() {
			close(msgQueue[p])
		}
	}()

	// custom msg
	go func() {
		for {
			select {
			case m, ok := <-msgQueue["HIGH"]:
				if !ok {
					msgQueue["HIGH"] = nil
					continue
				}
				fmt.Println("get high msg:", m)
			default:
				select {
				case m, ok := <-msgQueue["LOW"]:
					if !ok {
						msgQueue["LOW"] = nil
						continue
					}
					fmt.Println("get low msg:", m)
				default:
					fmt.Println("get nothing, sleep...")
				}
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
}
