/*Testing benchmark for http + gobs

Author: xpzouying@gmail.com

We have Student type for testing.
Provide http service, request message body is gobs body.
Marshal and unmarshal request message body with gobs.

*/

package main

import (
	"encoding/gob"
	"log"
	"net/http"
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
	dec := gob.NewDecoder(r.Body)

	var s Student
	if err := dec.Decode(&s); err != nil {
		log.Printf("decode request body error: %v", err)
		return
	}

	resp := Response{true, "finish"}
	enc := gob.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")
	if err := enc.Encode(resp); err != nil {
		log.Printf("encode response error: %v", err)
	}
}

func main() {
	http.HandleFunc("/api1", API1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
