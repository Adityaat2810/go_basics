package main

import (
	"fmt"
	"sync"
)

/*
------------ Mutex = Mutual Exclusion ------------

A mutex allows only one goroutine at a time to access a
critical section (shared data).

Why Mutex ?
    The real enemy: Race Condition
	    - Multiple goroutines read/write shared memory
        - At least one of them writes
        - No synchronization



*/

// No mutex

var countNoMutex int

func incrementNoMutex(wg *sync.WaitGroup){
  defer wg.Done()
  countNoMutex++
}

func main(){
  var wg sync.WaitGroup
  for i := 0; i<1000; i++{
	wg.Add(1)
	go incrementNoMutex(&wg)
  }

  wg.Wait()
  fmt.Println(countNoMutex)  // giveing 959 , 969 etc
  // Two goroutines can interleave → lost updates.

  /*

  =============== Enter Mutex ==================
  CORE IDEA
    - lock before accessing shared data
	- unlock after done
	- only one goroutine holds the lock

  BASIC MUTEX SYNTAX
    var mu sync.MUTEX

	mu.Lock()
  	// critical section
	mu.Unlock()

  */
  
}