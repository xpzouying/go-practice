# gRPC vs HTTP RESTFul

In these example, we will compare web service in these part,

1. message encode/decode
1. gRPC vs. HTTP RESTFul

I take these comparation,

- Message encode/decode comparation: json vs. gob vs. msgpack vs. protobuf
- gRPC vs. HTTP RESTFul

## Message encode/decode

To test code in golang, we have two way.

1. normal unittest
1. benchmark test

For normal unittest, we just run
```bash
go test -v
```
Go testing not run benchmark in default,

We need run the testcase with `-test.bench`.

```bash
go test -test.bench=".*"
```
OR
```bash
go test -bench=.
```
The benchmark result is following,

### HTTP + JSON

```bash
go test -v -bench=.
=== RUN   TestAPI1
--- PASS: TestAPI1 (0.00s)
goos: darwin
goarch: amd64
BenchmarkAPI1-8           200000              7643 ns/op
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/go-practice/gRPC_vs_HTTP_REST/http_protobuf_vs_http_json/http_json 1.638s
```

The testcase cost `1.638s`.
The loop ran `200000` times at speed of `7643 ns per loop`.

### HTTP + GOB

```bash
go test -v -bench=.
=== RUN   TestAPI1
--- PASS: TestAPI1 (0.00s)
goos: darwin
goarch: amd64
BenchmarkAPI1-8            20000             60853 ns/op
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/go-practice/gRPC_vs_HTTP_REST/http_protobuf_vs_http_json/http_gobs       1.886s
```
The testcase cost `1.886s`.
The loop ran `20000` times at speed of `60853 ns per loop`.

In our testcase gob is slower than json message.

However, in benchmark test at [go_serialization_benchmarks on github.com](https://github.com/alecthomas/go_serialization_benchmarks), gob is faster than json in marshal and unmarshal process.

In [msgpack github.com](https://github.com/vmihailenco/msgpack), testing is similar with my testing which json faster than gob.

### HTTP + msgpack

In this testcase, we test third marshal/unmarshal solution. We use msgpack tool to process message.

[https://github.com/vmihailenco/msgpack](https://github.com/vmihailenco/msgpack)

```bash
go test -v -bench .
=== RUN   TestAPI1
--- PASS: TestAPI1 (0.00s)
goos: darwin
goarch: amd64
BenchmarkAPI1-8           200000              7389 ns/op
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/go-practice/gRPC_vs_HTTP_REST/http_protobuf_vs_http_json/http_msgpack    1.579s
```

The result of testing is similar with json solution.

It's cost 7389 ns per operation.

### HTTP + protobuf

Using protobuf to encode/decode message, we need preprocess our type.

I create a file named `student.proto` and define our type in proto3.

Two type have defined.

1. Student: the information of student
1. Response: the response for request

```bash
go test -v -bench .
=== RUN   TestAPI1
--- PASS: TestAPI1 (0.00s)
goos: darwin
goarch: amd64
BenchmarkAPI1-8           300000              4755 ns/op
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/go-practice/gRPC_vs_HTTP_REST/http_protobuf_vs_http_json/http_protobuf   2.489s
```

The result is the best in our test.
