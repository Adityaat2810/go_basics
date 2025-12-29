package main

import (
	"fmt"
	"sync"
)


/**

8️⃣ Why Fan-In?
Collect results
Aggregate responses
Combine parallel computations


**/
func worker(id int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- id * 10
}

func main() {
	results := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Result:", res)
	}
}


/**

Fan-out = parallelism
Fan-in = coordination

**/