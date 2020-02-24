package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port   = flag.String("port", ":18080", "listen port for server")
	lbPort = flag.String("lb", ":9090", "register port in load balance")
)

func heartbeat(svr string) error {
	reqBody := []byte(*port)

	resp, err := http.Post(svr, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		log.Printf("ERROR: post")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not ok: %v", resp.StatusCode)
	}

	return nil
}

// send heart-beat
func registerService(svr string) error {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			if err := heartbeat(svr); err != nil {
				log.Printf("ERROR: heartbeat error: %v", err)
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()

	go registerService("http://localhost" + *lbPort + "/register")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("in server:%v", *port)

		w.Write([]byte(fmt.Sprintf("server-%s", *port)))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Panic(http.ListenAndServe(*port, nil))
}
