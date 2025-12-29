package main

import "fmt"

// Channel = a typed conduit for communication between goroutines
/*

Goroutine A(send)  ---- channel ---->  Goroutine B(recieve)

// syntax
ch := make(chan int)

- can send data
- can recieve data
- is thread safe
- provides synchronization

Basically, a channel allows goroutines to commucnicate and synchronize
without explicit lock


*/

//Why do we need channel ?

var x int

/*
func main(){
	var wt sync.WaitGroup
	wt.Add(1)
	go func(){
		wt.Done()
		x = 10
	}()

	wt.Add(1)
	go func() {
		wt.Done()
	    fmt.Println(x)
	}()

	// here race condition is happening beteen both goroutines -> sometimes print 10 , sometimes 0
	wt.Wait()

		// ⚠️ Data race
		// ⚠️ No ordering
		// ⚠️ Hard to reason

	fmt.Println("HEY DONE !")
}

*/

// solution

func main(){
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

    fmt.Println(<-ch)

	// no race
	// gaurnteed order
	// safe

	// [IMP!!] GO Philosopy
	/*
		DO not communicate by sharing memory; share memory by
		comminicating !

	*/

}