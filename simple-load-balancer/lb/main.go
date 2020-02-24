package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	// save all backend server list,
	// key: backend name
	// value: backend last heartbeat time
	backends = make(map[string]time.Time, 2)
)

func checkBackendsStatus() {
	for {
		select {
		case <-time.Tick(time.Second):
			for k, v := range backends {
				if time.Now().Sub(v) > 3*time.Second {
					log.Printf("WARN: heart-beat timeout, remove backend: %s", k)
					delete(backends, k)
				}
			}
		}
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: read request failed")
		return
	}

	log.Printf("server register:%s", body)

	// save backend to server list
	backendPort := string(body)
	backends[backendPort] = time.Now()

	w.Write([]byte("OK"))
}

func handleLB(w http.ResponseWriter, r *http.Request) {
	// choose the rand backend
	n := rand.Intn(2)
	var backend string
	for port := range backends {
		backend = port
		if n == 0 {
			break
		}
	}

	// send redirect request
	url := "http://localhost" + backend
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	port := flag.String("port", ":9090", "listen port")
	flag.Parse()

	go checkBackendsStatus()

	http.HandleFunc("/register", handleRegister)

	http.HandleFunc("/", handleLB)

	log.Printf("listen port: %s", *port)
	http.ListenAndServe(*port, nil)
}
