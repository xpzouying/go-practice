# when did goroutine switch

## endless_for.go

Goroutine didn't release OS threads and couldn't switch to other goroutine.

The programming print out the count of `id: ` equal `runtime.GOMAXPROCS(0)`.

## release_by_sleep.go

Release by `time.Sleep()`.

The code will print out all `id: `.

## release_by_gosched.go

Release by `runtime.Gosched()`

The code will print out all `id: `.