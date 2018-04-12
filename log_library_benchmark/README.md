# benchmark for logging

## run

### run by main.go

```bash
go run main.go
```

### run by benchmark

The lines of logging file is not equal to b.N. All logging library save the duplication logging when using benchmark.

```bash
go test -bench=.
```

## some tips

Two type of `io.Writer` for Out in logrus. The first one is write to Out directly. The second one is adding the cache for Out, like `bufio.Writer`.

We could see the result for this two benchmark, performance have big improvement: 13.465µs vs 4.81µs for per operation.

## result

logrus count:1000000, time_used: 13.465060626s, time_per_op: 13.465µs
logrusOutputBufio() - logrus count:1000000, time_used: 4.810216561s, time_per_op: 4.81µs
logrus parallel count:1000000, time_used: 16.971633791s, time_per_op: 16.971µs
zerolog count:1000000, time_used: 9.131216964s, time_per_op: 9.131µs
uber zap count:1000000, time_used: 13.219052873s, time_per_op: 13.219µs
stdlog count:1000000, time_used: 9.053266913s, time_per_op: 9.053µs