package main

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI1(t *testing.T) {
	s := Student{
		Name:        "zouying",
		Age:         32,
		Description: "here is self description",
	}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(s); err != nil {
		t.Fatalf("encode a new student error: %v", err)
	}

	r := httptest.NewRequest(http.MethodPost, "http://this.is.test", buf)
	w := httptest.NewRecorder()

	API1(w, r)

	if w.Code != http.StatusOK {
		t.Fatal("status code != 200")
	}

	dec := gob.NewDecoder(w.Body)
	var resp Response
	if err := dec.Decode(&resp); err != nil {
		t.Fatalf("decode response error: %v", err)
	}
}

func BenchmarkAPI1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Student{
			Name:        "zouying",
			Age:         32,
			Description: "here is self description",
		}

		buf := new(bytes.Buffer)
		enc := gob.NewEncoder(buf)
		if err := enc.Encode(s); err != nil {
			b.Fatalf("encode a new student error: %v", err)
		}

		r := httptest.NewRequest(http.MethodPost, "http://this.is.test", buf)
		w := httptest.NewRecorder()

		API1(w, r)

		if w.Code != http.StatusOK {
			b.Fatal("status code != 200")
		}

		dec := gob.NewDecoder(w.Body)
		var resp Response
		if err := dec.Decode(&resp); err != nil {
			b.Fatalf("decode response error: %v", err)
		}
	}
}
