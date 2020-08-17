package main

const nRequester = 100 // Number of entities requesting work
const nWorker = 10     // Number of available workers

func main() {
	work := make(chan Request)
	for i := 0; i < nRequester; i++ {
		go createAndRequest(work)
	}
	NewBalancer().balance(work)
}
