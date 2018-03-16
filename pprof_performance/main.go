package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
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
	f, err := os.Create("./cpuprofile")
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(f)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		select {
		case v := <-c:
			fmt.Printf("get signal %v\n", v)
			pprof.StopCPUProfile()
			return

		default:
			getWeb()
		}
	}

}
