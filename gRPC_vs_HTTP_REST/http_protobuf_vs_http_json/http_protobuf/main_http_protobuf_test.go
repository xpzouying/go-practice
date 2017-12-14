package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	proto "github.com/golang/protobuf/proto"
)

func TestAPI1(t *testing.T) {

	s := Student{
		Name:        "zouying",
		Age:         32,
		Description: "here is self description",
	}

	reqBody, err := proto.Marshal(&s)
	if err != nil {
		t.Fatalf("marshal request body error: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "http://this.is.test", bytes.NewReader(reqBody))
	w := httptest.NewRecorder()

	API1(w, req)

	var resp Response
	if err := proto.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("msgpack decode error: %v", err)
	}

	if resp.Success != true {
		t.Fatal("response error")
	}
}

func BenchmarkAPI1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Student{
			Name:        "zouying",
			Age:         32,
			Description: "here is self description",
		}

		reqBody, err := proto.Marshal(&s)
		if err != nil {
			b.Fatalf("marshal request body error: %v", err)
		}

		req := httptest.NewRequest(http.MethodPost, "http://this.is.test", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		API1(w, req)

		var resp Response
		if err := proto.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			b.Fatalf("msgpack decode error: %v", err)
		}

		if resp.Success != true {
			b.Fatal("response error")
		}
	}
}
