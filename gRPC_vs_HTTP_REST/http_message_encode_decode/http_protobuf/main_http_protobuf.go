/*Testing benchmark for http + msgpack

Author: xpzouying@gmail.com

We have Student type for testing.
Provide http service, request message body is msgpack body.
Marshal and unmarshal request message body with msgpack.

*/

package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// API1 handle /api1 request
func API1(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read request body error: %v", err)
		return
	}

	var s Student
	if err := proto.Unmarshal(reqBody, &s); err != nil {
		log.Printf("unmarshal request body by proto error: %v", err)
		return
	}

	resp := Response{
		Success: true,
		Message: "finish",
	}
	body, err := proto.Marshal(&resp)
	if err != nil {
		log.Printf("proto marshal result error: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	http.HandleFunc("/api1", API1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
