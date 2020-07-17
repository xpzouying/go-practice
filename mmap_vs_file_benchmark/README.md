# MMap vs os.File Performance

Compare performance between mmap and os.File in Go.

## Run

```bash
go test -bench=. -v . -benchmem
```

```
goos: darwin
goarch: amd64
BenchmarkMMap
BenchmarkMMap-4            42883             25398 ns/op             360 B/op          5 allocs/op
BenchmarkFileOpen
BenchmarkFileOpen-4        82740             15660 ns/op             120 B/op          3 allocs/op
PASS
ok      _/Users/zy/src/GOPATH/src/github.com/xpzouying/go-practice/mmap_vs_file_benchmark       4.853s
```

## mmap

In the computing, [mmap (2)](https://en.wikipedia.org/wiki/Mmap) is a POSIX-compliant Unix system call that maps files or devices into memory. It is a method of memory-mapped file I/O.

- Use [mmap package](golang.org/x/exp/mmap)


Open a file by: `mmap.Open("./README.md")`

## os.File

Open a file by: `os.Open("./README.md")`
