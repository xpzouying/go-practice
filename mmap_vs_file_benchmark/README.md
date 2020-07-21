# MMap vs os.File Performance

Compare performance between mmap and os.File in Go.

## Run

```bash
go test -bench=. -v . -benchmem
```

```
goos: darwin
goarch: amd64
pkg: github.com/xpzouying/go-practice/mmap_vs_file_benchmark
BenchmarkMMap
BenchmarkMMap-4       	153527704	         7.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkFileOpen
BenchmarkFileOpen-4   	 1232979	       983 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xpzouying/go-practice/mmap_vs_file_benchmark	4.582s
```

## mmap

In the computing, [mmap (2)](https://en.wikipedia.org/wiki/Mmap) is a POSIX-compliant Unix system call that maps files or devices into memory. It is a method of memory-mapped file I/O.

- Use [mmap package](golang.org/x/exp/mmap)


Open a file by: `mmap.Open("./README.md")`

## os.File

Open a file by: `os.Open("./README.md")`
