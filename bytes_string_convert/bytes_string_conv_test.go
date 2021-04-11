package main

import (
	"reflect"
	"testing"
)

var (
	strdata  = "github.com/xpzouying/go-practice"
	bytedata = []byte(strdata)

	b2sFunc = map[string]func([]byte) string{
		"string(bytes) func:": normalb2s,
		"b2s-use pointer:":    b2s,
	}

	s2bFunc = map[string]func(string) []byte{
		"[]byte(s) func:":  normals2b,
		"s2b-use pointer:": s2b,
	}
)

func normalb2s(b []byte) string { return string(b) }

func normals2b(s string) []byte { return []byte(s) }

func Test_b2s(t *testing.T) {
	got1, got2 := normalb2s(bytedata), b2s(bytedata)
	if !reflect.DeepEqual(got1, got2) {
		t.Errorf("b2s error, not equal. %s != %s", got1, got2)
	}
}

func Benchmark_b2s(b *testing.B) {

	for name, convfunc := range b2sFunc {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {

				_ = convfunc(bytedata)
			}
		})
	}
}

func Benchmark_s2b(b *testing.B) {

	for name, convfunc := range s2bFunc {
		b.Run(name, func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				_ = convfunc(strdata)
			}

		})
	}

}
