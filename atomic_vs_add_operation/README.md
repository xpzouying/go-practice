# compare sync/atomic vs add operation

Why golang have [AddInt32](https://golang.org/pkg/sync/atomic/#AddInt32), not only use `i += num` model?

Goroutine/Coroutine is one of best features in golang, we always over use goroutine.

If multiple goroutine write the shared memory, the error will happened.

Run the programming, we have `counter` for counting adding.

```bash
go run add_op.go
# counter: {count:1815}

go run add_op.go
# counter: {count:1282}
```

We add 2000 times, but result is wrong.

Why?

Because we use two goroutine to write the same memory (count in counter).

How to solve this?

1. Use mutex.

```bash
go run add_op_with_mutex.go
# counter: {count:2000 mu:{state:0 sema:0}}
```

2. Use `atomic.AddInt32/64`


```bash
go run atomic.go
# counter: {count:2000}
```

Benchmark

Compare mutex increase and atomic increase.

```bash
bash ./test_benchmark.sh
benchkmark mutex
goos: darwin
goarch: amd64
BenchmarkMutexCounter-4          3000000               467 ns/op
PASS
ok      command-line-arguments  1.897s
benchkmark atomic
goos: darwin
goarch: amd64
BenchmarkAtomicCounter-4         3000000               487 ns/op
PASS
ok      command-line-arguments  1.954s
```