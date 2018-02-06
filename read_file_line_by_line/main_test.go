package main

import "testing"

func BenchmarkScanFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := scanFile(); err != nil {
			b.Fatalf("scan file error: %v", err)
		}
	}
}

func BenchmarkReadFileLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := readFileLines(); err != nil {
			b.Fatalf("scan file error: %v", err)
		}
	}
}

func BenchmarkReadFileOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := readFileOnce(); err != nil {
			b.Fatalf("scan file error: %v", err)
		}
	}
}
