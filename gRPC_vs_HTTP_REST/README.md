# gRPC service

The gRPC is more faster than HTTP service.

When we use `http+json` service, we need to encode or decode request and response message.

As my benchmark testing show, only `4.39 ns` per operation. In contrast, HTTP+JSON cost `7643 ns` per operation.

```bash
go test -v -bench .
=== RUN   TestAPI1
--- PASS: TestAPI1 (0.00s)
goos: darwin
goarch: amd64
BenchmarkAPI1-8         300000000                4.39 ns/op
PASS
ok      _/Users/zouying/src/Github.com/ZOUYING/go-practice/gRPC_vs_HTTP_REST/grpc_protobuf      1.799s
```