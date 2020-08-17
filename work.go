package main

import "math"

type Work struct {
	idx      int          // heap index
	work     chan Request // channel for this worker
	pending  int          // number of pending requests for this worker
}

// Do work indefinitely by pulling the next available request from the work channel.
// The response is written to the response channel and we complete the work being
// done by writting to the done channel.
//
func (w *Work) doWork(done chan *Work) {
	// worker works indefinitely
	for {
		// extract request from WOK channel
		req := <-w.work
		// write to RESP channel
		req.resp <- math.Sin(float64(req.data))
		// write to DONE channel
		done <- w
	}
}
