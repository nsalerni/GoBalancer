package main

import (
	"math/rand"
	"time"
)

type Request struct {
	data int
	resp chan float64
}

// For demo purposes we spawn requests indefinitely, waiting a random number of
// milliseconds before sending the next request.
//
func createAndRequest(req chan Request) {
	resp := make(chan float64)

	for {
		time.Sleep(time.Duration(rand.Int63n(int64(time.Millisecond))))
		req <- Request{int(rand.Int31n(90)), resp}
		<-resp
	}
}