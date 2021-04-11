# bytes string之间的快速转换

看到`fasthttp`中有将string和bytes快速转换的辅助函数。[链接点击](https://github.com/valyala/fasthttp/blob/c48d3735fa9864a7c1724168812f3571c8313581/bytesconv.go#L387)

```go
// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// s2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}
```

测试一下和普通性能对比。

测试`[]byte`转换`string`的性能。

```bash
>> go test -bench=. -run=none -count=10 .
goos: darwin
goarch: arm64
pkg: github.com/xpzouying/go-practice/bytes_string_convert
Benchmark_b2s/string(bytes)_func:-8             75367821                15.60 ns/op
Benchmark_b2s/string(bytes)_func:-8             76076361                15.25 ns/op
Benchmark_b2s/string(bytes)_func:-8             75109142                15.24 ns/op
Benchmark_b2s/string(bytes)_func:-8             76953082                15.25 ns/op
Benchmark_b2s/string(bytes)_func:-8             75058249                15.26 ns/op
Benchmark_b2s/string(bytes)_func:-8             75734664                15.30 ns/op
Benchmark_b2s/string(bytes)_func:-8             76914856                15.32 ns/op
Benchmark_b2s/string(bytes)_func:-8             76271793                15.30 ns/op
Benchmark_b2s/string(bytes)_func:-8             76444280                15.44 ns/op
Benchmark_b2s/string(bytes)_func:-8             78698845                15.31 ns/op
Benchmark_b2s/b2s-use_pointer:-8                812254438                1.474 ns/op
Benchmark_b2s/b2s-use_pointer:-8                816100992                2.074 ns/op
Benchmark_b2s/b2s-use_pointer:-8                815082879                1.470 ns/op
Benchmark_b2s/b2s-use_pointer:-8                815812945                2.076 ns/op
Benchmark_b2s/b2s-use_pointer:-8                816684404                1.471 ns/op
Benchmark_b2s/b2s-use_pointer:-8                816206455                1.471 ns/op
Benchmark_b2s/b2s-use_pointer:-8                817465142                1.470 ns/op
Benchmark_b2s/b2s-use_pointer:-8                815514482                1.471 ns/op
Benchmark_b2s/b2s-use_pointer:-8                815448441                1.472 ns/op
Benchmark_b2s/b2s-use_pointer:-8                816148864                1.470 ns/op
PASS
ok      github.com/xpzouying/go-practice/bytes_string_convert   28.396s
```

测试`string`转换为`[]byte`的性能，

```bash
>> go test -bench=Benchmark_s2b -run=none -count=5 .
goos: darwin
goarch: arm64
pkg: github.com/xpzouying/go-practice/bytes_string_convert
Benchmark_s2b/s2b-use_pointer:-8                517091395                2.069 ns/op
Benchmark_s2b/s2b-use_pointer:-8                575813779                2.067 ns/op
Benchmark_s2b/s2b-use_pointer:-8                577440252                2.068 ns/op
Benchmark_s2b/s2b-use_pointer:-8                578506900                2.071 ns/op
Benchmark_s2b/s2b-use_pointer:-8                576001612                2.070 ns/op
Benchmark_s2b/[]byte(s)_func:-8                 56602884                20.28 ns/op
Benchmark_s2b/[]byte(s)_func:-8                 57922862                20.29 ns/op
Benchmark_s2b/[]byte(s)_func:-8                 58262206                20.22 ns/op
Benchmark_s2b/[]byte(s)_func:-8                 58290153                20.34 ns/op
Benchmark_s2b/[]byte(s)_func:-8                 58049416                20.48 ns/op
PASS
ok      github.com/xpzouying/go-practice/bytes_string_convert   13.100s
```