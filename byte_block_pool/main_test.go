package bytespool

import (
	"fmt"
	"testing"
)

func TestBytesPoolOnce(t *testing.T) {
	const (
		maxElementSize = 32
		maxElementNum  = 10
	)

	bp := NewBytesPool(maxElementSize, maxElementNum)

	data := "test_data"

	if err := bp.Set(0, []byte(data)); err != nil {
		t.Fatalf("bytes pool set data error: %v", err)
	}

	got, err := bp.Get(0)
	if err != nil {
		t.Fatalf("bytes pool get data error: %v", err)
	}

	if string(got) != data {
		t.Fatalf("got bytes from pool not expect")
	}
}

func TestBytesPool(t *testing.T) {
	const (
		maxElementSize = 32
		maxElementNum  = 10
	)

	bp := NewBytesPool(maxElementSize, maxElementNum)

	for i := 0; i < maxElementNum; i++ {
		data := fmt.Sprintf("num%d", i)

		if err := bp.Set(i, []byte(data)); err != nil {
			t.Fatalf("bytes pool set data error: %v", err)
		}

		got, err := bp.Get(i)
		if err != nil {
			t.Fatalf("bytes pool get data error: %v", err)
		}

		if string(got) != data {
			t.Fatalf("got bytes from pool not expect")
		}
	}
}
