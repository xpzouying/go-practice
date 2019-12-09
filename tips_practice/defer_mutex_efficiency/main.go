package main

import (
	"os"
)

func closeFileDirectly() error {
	for i := 0; i < 1000; i++ {
		f, err := os.OpenFile("main.go", os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		// ...
		f.Close()
	}

	return nil
}

func closeFileByDefer() error {
	for i := 0; i < 1000; i++ {
		f, err := os.OpenFile("main.go", os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer f.Close()
		// ...
	}

	return nil
}

func main() {

}
