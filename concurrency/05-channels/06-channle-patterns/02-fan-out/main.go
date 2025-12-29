package main

import (
	"fmt"
	"sync"
)

/**
Fan-out = one producer → many workers

- You distribute work from a single source (channel) to multiple goroutines so tasks run in parallel.
“Same input channel, multiple consumers.”
**/

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func main(){
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// fan-out 3 workers
	for i:=0; i<=3; i++{
	  wg.Add(1)
	  go worker(i, jobs, &wg)
	}

	// send
	for j := 1; j <= 6; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()
}