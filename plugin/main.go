package main

import (
	"log"
	"plugin"
)

func main() {

	p, err := plugin.Open("./lower_plugin.so")
	panicError(err)

	lower, err := p.Lookup("Lower")
	panicError(err)

	lowerFunc, ok := lower.(func(string) string)
	if !ok {
		log.Fatalf("unexpected type from module symbol Lower: %T", lower)
	}

	got := lowerFunc("HELLO")

	if got != "hello" {
		log.Fatalf("unexpected result from Lower: %q", got)
	} else {
		log.Printf("got: %q", got)
	}
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}
