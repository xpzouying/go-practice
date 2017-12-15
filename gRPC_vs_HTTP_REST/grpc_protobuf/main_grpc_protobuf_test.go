package main

import (
	"testing"

	"context"
)

func TestAPI1(t *testing.T) {
	s := server{}

	stu := Student{
		Name:        "zouying",
		Age:         32,
		Description: "this is student zouying",
	}

	resp, err := s.API1(context.Background(), &stu)
	if err != nil {
		t.Fatalf("service API1 error: %v", err)
	}

	if resp.Success != true {
		t.Fatalf("response result should be true, but got not true")
	}
}

func BenchmarkAPI1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server{}

		stu := Student{
			Name:        "zouying",
			Age:         32,
			Description: "this is student zouying",
		}

		resp, err := s.API1(context.Background(), &stu)
		if err != nil {
			b.Fatalf("service API1 error: %v", err)
		}

		if resp.Success != true {
			b.Fatalf("response result should be true, but got not true")
		}
	}
}
