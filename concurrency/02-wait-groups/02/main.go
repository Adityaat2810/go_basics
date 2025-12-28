package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup){
  defer wg.Done()
  fmt.Println("Worker ", id, " done")
}

func main(){
  var wg sync.WaitGroup

  for i := 0; i<3; i++{
	wg.Add(1)
	go worker(i, &wg)
  }

  wg.Wait()
  fmt.Println("ALL workers finished !")

}

/*
* Very important WaitGroup rules (INTERVIEW TRAPS)

1. Add() before goroutine starts

	wrong
	go func() {
		wg.Add(1)   // race condition
		defer wg.Done()
	}()

	right
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	why
	Wait() may run before Add()
	Causes panic or undefined behavior

2. Always pass WaitGroup by pointer
   (Copying WaitGroup copies its counter → broken sync)

3. Done() must be called exactly once per Add(1)
	Missing Done() → deadlock
	Extra Done() → panic (negative counter)


*/