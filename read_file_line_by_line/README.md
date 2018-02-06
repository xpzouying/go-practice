# Read file line by line

Run test:

```bash
# change const logfile = "logfile.log"
go run main.go
```

Test file:

- size: ~250M
- lines: 2,500,000 lines

Three ways read file line by line,

1. Use bufio.NewScanner
2. Use ReadString in bufio.Reader
3. Use ioutil.ReadFile, and strings.Split

Result (in Macbook Pro 2016):

1. bufio.NewScanner: scan file, time_used: 0.383570461, lines=2493999
2. bufio.Reader: reader read string in file, time_used: 0.465576298, lines=2493999
3. ioutil.ReadFile and strings.Split: read file once and split strings, time_used: 0.424787524, lines=2494000
>Use strings.Split() to split string, will get one more line, the last string is empty("").