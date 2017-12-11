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
