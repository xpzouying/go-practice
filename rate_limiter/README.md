# rate limiter

使用Golang提供的`golang.org/x/time/rate`进行并发限制。


## URLs

- [Github Repo](https://pkg.go.dev/mod/golang.org/x/time)
- [Docs](https://pkg.go.dev/golang.org/x/time/rate)


## 说明

**Run**

```bash
# run
go run main.go 2>&1  | tee 1.log

# check qps
grep "finish task:" 1.log | awk '{print $2}' | sort | uniq -c | sort -r
```