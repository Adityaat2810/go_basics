package main

import (
	"fmt"
	"sync"
)

var (
  count int
  mu sync.Mutex

)

func increment(wg *sync.WaitGroup){
  defer wg.Done()

  mu.Lock()
  count++
  mu.Unlock()  // Only protect what is necessary, not more.

  // best practice
  // defer + Mutex (very important)
  /**
	mu.Lock()
    count++
    defer mu.Unlock()
  */

  // Why?
  // 1. Prevents deadlocks
  // 2. Unlock happens even if function returns early or panics

}

func main(){
  var wg sync.WaitGroup
  for i:=0 ; i<1000; i++{
	wg.Add(1)
	go increment(&wg)
  }

  wg.Wait()
  fmt.Println("okay all done ", count)

}
