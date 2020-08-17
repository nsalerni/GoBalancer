package main

import (
	"container/heap"
	"fmt"
)

type Balancer struct {
	pool Pool       // a pool of workers
	done chan *Work // a done channel to notify of completed work
}

// A new balancer with a pool of workers and a done channel
//
func NewBalancer() *Balancer {
	b := &Balancer{make(Pool, 0, nWorker), make(chan *Work, nWorker)}
	b = b.sampleRequests()
	return b
}

// Sample requests for the load balancer, for illustrative purposes.
// In reality, this will be replaced with real requests
//
func (b *Balancer) sampleRequests() *Balancer {
	for i := 0; i < nWorker; i++ {
		w := &Work{work: make(chan Request, nRequester)}
		heap.Push(&b.pool, w)
		go w.doWork(b.done)
	}
	return b
}

// Balance the work from the request channel
//
func (b *Balancer) balance(req chan Request) {
	for {
		select {
		// process requests from request channel
		case request := <-req:
			b.dispatch(request)
		// read from done channel
		case w := <-b.done:
			b.completed(w)
		}
		b.print()
	}
}

// Dispatch request to the least loaded worker.
// This implementation of a load balancer uses the "least busy" scheduling strategy
// that schedules jobs to the least busy worker based on their current workload.
//
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker
	w := heap.Pop(&b.pool).(*Work)
	w.work <- req
	w.pending++
	// Update the heap while the request is being processed
	heap.Push(&b.pool, w)
}

// Mark work as complete and place worker back on the heap (indicating available for work)
//
func (b *Balancer) completed(w *Work) {
	w.pending--
	heap.Remove(&b.pool, w.idx)
	heap.Push(&b.pool, w)
}

// Show the status of each worker, in addition to the following stats:
//    avg: average load across the worker pool
//    std: variance of the load across worker pool (shows how well work is distributed, lower + stable is better)
//
func (b *Balancer) print() {
	sum := 0
	sumsq := 0

	for _, w := range b.pool {
		fmt.Printf("%d ", w.pending)
		sum += w.pending
		sumsq += w.pending * w.pending
	}

	avg := float64(sum) / float64(len(b.pool))
	variance := float64(sumsq)/float64(len(b.pool)) - avg*avg
	fmt.Printf(" %.2f %.2f\n", avg, variance)
}