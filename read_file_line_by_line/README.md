# Read file line by line

Run test:

```bash
# change const logfile = "logfile.log"
go run main.go
```

Test file:

- size: ~250M
- lines: 2,500,000 lines

Two ways read file line by line,

1. Use bufio.NewScanner
2. Use ReadString in bufio.Reader

Result (in Macbook Pro 2016):

1. bufio.NewScanner: scan file, time_used: 0.395491384, lines=2493999
2. bufio.Reader: reader read string in file, time_used: 0.446867622, lines=2493999