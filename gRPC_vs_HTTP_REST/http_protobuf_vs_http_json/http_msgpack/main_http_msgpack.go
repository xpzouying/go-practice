/*Testing benchmark for http + msgpack

Author: xpzouying@gmail.com

We have Student type for testing.
Provide http service, request message body is msgpack body.
Marshal and unmarshal request message body with msgpack.

*/

package main

import (
	"log"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

// Student is common type for testing marshal and unmarshal
type Student struct {
	Name        string
	Age         int
	Description string
}

// Response is response for request
type Response struct {
	Success bool
	Message string
}

// API1 handle /api1 request
func API1(w http.ResponseWriter, r *http.Request) {
	var s Student
	if err := msgpack.NewDecoder(r.Body).Decode(&s); err != nil {
		log.Printf("decode request body error: %v", err)
		return
	}

	resp := Response{true, "finish"}
	body, err := msgpack.Marshal(resp)
	if err != nil {
		log.Printf("msgpack marshal result error: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	http.HandleFunc("/api1", API1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
