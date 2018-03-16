package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func getWeb() {
	resp, err := http.Get("https://stackoverflow.com/questions/35860642/golang-how-to-get-computing-time-using-pprof-within-a-web-server")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("body: %s\n", body)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		getWeb()
	}

}
