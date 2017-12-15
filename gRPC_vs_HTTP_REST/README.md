# Message encode and decode in HTTP service

In these example, we will compare message encode and decode in HTTP web service.

1. message encode/decode (Part. I)
    - gob
    - json
    - msgpack
    - protobuf
1. gRPC vs. HTTP RESTFul (Part. II)
    - gRPC
    - HTTP service

## Result in testing

### TL;DR

| Encode/Decode Tools | Count of loop | Time used all | Time used for single |
|---|---|---|---|
| JSON | 200000 | 1.638s | 7643 ns/op |
| GOB | 20000 | 1.886s | 60853 ns/op |
| msgpack | 200000 | 1.579s | 7389 ns/op |
| protobuf | 300000 | 2.489s | 4755 ns/op |

In my testing, `protobuf` is the fastest, but need preprocess. 1. create proto file; 2. compile protfo file to target language.

`protobuf` is also not for human reading, if you want read the message after encoded, like in browser display, `json` is a better choose.

## Message encode/decode

Two test way in golang,

1. normal unittest: `go test`
1. benchmark test: `go test -bench`

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
