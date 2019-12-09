package main

import (
	"fmt"
	"testing"
)

func TestShard(t *testing.T) {
	s := newShard()

	const COUNT = 10
	m := make(map[string]string, COUNT)
	for i := 0; i < COUNT; i++ {
		key := fmt.Sprintf("key-%d", i)
		val := fmt.Sprintf("value-%d", i)
		m[key] = val
	}

	// setting
	for k, v := range m {
		s.set([]byte(k), []byte(v))
	}

	// verify
	for k, v := range m {
		got, err := s.get([]byte(k))
		if err != nil {
			t.Errorf("cache shard get error: %v", err)
			continue
		}

		if v != string(got) {
			t.Errorf("key=%s,value=%s, but got=%s", k, v, got)
		}
	}
}
