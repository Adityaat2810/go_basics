package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup){
  defer wg.Done() // // called when one goroutine finishes
  // counter --
  fmt.Println("namastey")
}

func main(){

	// go sayHello()
	// go sayHello()
	// go sayHello()

	/* The Problem
	*    - main() exit
	*    - goroutine get killed
	*    - worker may never finished
	*  What we need:
    *  “Wait until all goroutines finish.”
	*/

	/// SOlution
	// sync.WaitGroup is a synchronization primitive used to
	// wait for a collection of goroutines to finish
	// It does ONLY ONE JOB: waiting.

	var wg sync.WaitGroup
	n := 3
	wg.Add(n)  // n = no of goroutines to wait for

	go sayHello(&wg)
	// [IMP] The goroutine that does the work must call Done()
	go sayHello(&wg)
	go sayHello(&wg)

	wg.Wait()  // blocks until counter becomes zero (all coroutine done)
	fmt.Println("all done")



}
