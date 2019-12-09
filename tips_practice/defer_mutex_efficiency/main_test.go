package main

import "testing"

func BenchmarkCloseFileDirectly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := closeFileDirectly(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCloseFileByDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := closeFileByDefer(); err != nil {
			b.Fatal(err)
		}
	}
}
