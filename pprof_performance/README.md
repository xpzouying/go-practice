# pprof performance

Measure cpuprofile for endless program.

1. create cpuprofile file
2. StartCPUProfile with this file
3. Create a channel to listen signal to accept <CTRL-C> signal

Usage:

1. go build .
2. run programming, `./pprof_performance`
3. generate svg for torch, `go-torch ./pprof_performance cpuprofile`
4. open svg with browser: `open torch.svg`