package main

import (
	"os"
	"testing"

	"golang.org/x/exp/mmap"
)

func BenchmarkMMap(b *testing.B) {
	buff := make([]byte, 32)

	for i := 0; i < b.N; i++ {
		at, err := mmap.Open("./README.md")
		if err != nil {
			b.Fatal(err)
		}

		if _, err := at.ReadAt(buff, 0); err != nil {
			b.Error(err)
		}

		at.Close()
	}
}

func BenchmarkFileOpen(b *testing.B) {
	buff := make([]byte, 32)

	for i := 0; i < b.N; i++ {
		f, err := os.Open("./README.md")
		if err != nil {
			b.Fatal(err)
		}

		if _, err := f.ReadAt(buff, 0); err != nil {
			b.Error(err)
		}

		f.Close()
	}
}
