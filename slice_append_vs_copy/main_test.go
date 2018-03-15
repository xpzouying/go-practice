package main

import (
	"fmt"
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	var s1 []string
	var s2 []string
	for i := 0; i < 10; i++ {
		s1 = append(s1, fmt.Sprintf("stringa-%d", i))
		s2 = append(s2, fmt.Sprintf("stringb-%d", i))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s1 = append(s1, s2...)
	}
	b.StopTimer()
}

func BenchmarkCopyStrings(b *testing.B) {
	var s1 []string
	var s2 []string
	for i := 0; i < 10; i++ {
		s1 = append(s1, fmt.Sprintf("stringa-%d", i))
		s2 = append(s2, fmt.Sprintf("stringb-%d", i))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s1 = copyString(s1, s2...)
	}
	b.StopTimer()
}

func copyString(slice []string, data ...string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
