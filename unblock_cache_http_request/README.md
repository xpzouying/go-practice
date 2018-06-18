# unblock cache http request

Ref: https://www.youtube.com/watch?v=KlDWmTcyXdA&t=651s&list=WL&index=40


Create a cache for http request without blocking.


**RUN**

`go test -v`


**RESULT**

```bash
go test -v
=== RUN   TestRequestAll
call request function:  http://www.weibo.com
call request function:  http://www.baidu.com
call request function:  http://www.sina.com.cn
call request function:  http://www.12306.cn
call request function:  http://www.163.com
call request function:  http://www.hao123.com
--- PASS: TestRequestAll (0.86s)
        main_test.go:56: request result: http://www.12306.cn,200 OK
        main_test.go:56: request result: http://www.baidu.com,200 OK
        main_test.go:56: request result: http://www.baidu.com,200 OK
        main_test.go:56: request result: http://www.hao123.com,200 OK
        main_test.go:56: request result: http://www.163.com,200 OK
        main_test.go:56: request result: http://www.weibo.com,200 OK
        main_test.go:56: request result: http://www.weibo.com,200 OK
        main_test.go:56: request result: http://www.weibo.com,200 OK
        main_test.go:56: request result: http://www.sina.com.cn,200 OK
        main_test.go:62: used_time: 855.036733ms
PASS
ok  _/go-practice/unblock_cache_http_request   0.874s
```