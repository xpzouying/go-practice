package main

import (
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		bytesToString()
	}
}

// --- BenchmarkBytesToString
// goos: darwin
// goarch: amd64
// pkg: github.com/xpzouying/go-practice/unsafe_pointer
// cpu: VirtualApple @ 2.50GHz
// BenchmarkBytesToString-8   	1000000000	         0.3327 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/xpzouying/go-practice/unsafe_pointer	0.539s

func BenchmarkBytesToString_2(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		data := []byte("zouying")

		name := string(data)
		// log.Printf("name=%s", name)
		_ = name
	}
}

// --- BenchmarkBytesToString_2
// goos: darwin
// goarch: amd64
// pkg: github.com/xpzouying/go-practice/unsafe_pointer
// cpu: VirtualApple @ 2.50GHz
// BenchmarkBytesToString_2-8   	158631984	         7.461 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/xpzouying/go-practice/unsafe_pointer	5.131s
