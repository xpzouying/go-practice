    # Middleware in web service (Golang)

    有时候我们需要统计web service接口的数据，比如记录日志、统计API调用时间、或者对HandleFunc进行错误处理，

    这个时候，middleware就很有帮助。

    ## 最初版本

    我们目前有个程序，监听8080端口，提供两个接口，
    
        1. compute：进行计算，耗时在500ms-1000ms之间
        2. version：得到版本号，耗时在10ms-60ms之间


    ```go
    package main

    import (
        "log"
        "math/rand"
        "net/http"
        "time"
    )

    // Compute is compute process for doing task
    func Compute(w http.ResponseWriter, r *http.Request) {
        time.Sleep(time.Duration((500 + rand.Intn(500))) * time.Millisecond)
    }


    // Version is getting version task
    func Version(w http.ResponseWriter, r *http.Request) {
        time.Sleep(time.Duration((10 + rand.Intn(50))) * time.Millisecond)
        w.Write([]byte("version"))
    }

    func main() {
        http.HandleFunc("/compute", Compute)
        http.HandleFunc("/version", Version)

        log.Panic(http.ListenAndServe(":8080", nil))
    }
    ```

    调用结果如下，

    ```bash
    # curl http://localhost:8080/compute
    finish
    ```

    ## 增加日志

    需要记录每次调用接口的名字。

    有两种方式记录日志，

        1. 在每一个HandleFunc第一行输出URI
        2. 增加中间件封装HandleFunc

    对于第一种情况，修改Compute()函数为，

    ```go
    // Compute is compute process for doing task
    func Compute(w http.ResponseWriter, r *http.Request) {
        log.Printf(r.RequestURI)

        time.Sleep(time.Duration((500 + rand.Intn(500))) * time.Millisecond)
        w.Write([]byte("finish"))
    }
    ```

    运行得到结果为，

    ```bash
    2017/12/11 23:07:24 /compute
    ```

    对于API少的情况，还比较适用。但是对于API接口比较多的情况，修改每一个函数就不太合适。并且当我们不光想统计URI信息时，还需要统计每一个调用接口的其他信息时，就需要修改每一处HandleFunc。

    所以我们可以使用`Middleware`形式来完成。

    我们希望使用middlware以后，对于main()函数变为，

    ```go
    func main() {
        http.HandleFunc("/compute", middleware(Compute))
        http.HandleFunc("/version", middleware(Version))

        log.Panic(http.ListenAndServe(":8080", nil))
    }
    ```

    使用middlware对每一个http handle func进行封装，在middleware中进行相应的需求处理，比如日志记录、错误处理、用时统计等。

    ```go
    func middleware(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
        return func(w http.ResponseWriter, r *http.Request) {
            begin := time.Now()
            defer func() {
                log.Printf("time_used: %v", time.Now().Sub(begin))
            }()
            // logging
            log.Printf(r.RequestURI)

            fn(w, r)
        }
    }
    ```

    `middlware`入参和返回值都为`func(w http.ResponseWriter, r *http.Request)`类型。

    在`middleware`中，进行了调用接口URI的统计和用时统计。

    完整的代码如下，

    ```go
    package main

    import (
        "log"
        "math/rand"
        "net/http"
        "time"
    )

    // Compute is compute process for doing task
    func Compute(w http.ResponseWriter, r *http.Request) {
        time.Sleep(time.Duration((500 + rand.Intn(500))) * time.Millisecond)
        w.Write([]byte("finish"))
    }

    // Version is getting version task
    func Version(w http.ResponseWriter, r *http.Request) {
        time.Sleep(time.Duration((10 + rand.Intn(50))) * time.Millisecond)
        w.Write([]byte("version"))
    }

    func middleware(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
        return func(w http.ResponseWriter, r *http.Request) {
            begin := time.Now()
            defer func() {
                log.Printf("time_used: %v", time.Now().Sub(begin))
            }()
            // logging
            log.Printf(r.RequestURI)

            fn(w, r)
        }
    }

    func main() {
        http.HandleFunc("/compute", middleware(Compute))
        http.HandleFunc("/version", middleware(Version))

        log.Panic(http.ListenAndServe(":8080", nil))
    }
    ```