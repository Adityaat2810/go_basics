package main

import (
	"fmt"
	"time"
)

func worker(ch chan int){
  time.Sleep(2* time.Second)
  fmt.Println("Recieved", <-ch)
}
func main(){
  ch := make(chan int, 1)

  go worker(ch)

  ch <- 42 // does NOT block
  fmt.Println("Sent 42")

  time.Sleep(3 * time.Second)

  /**
   op
   Sent 42
   Received: 42

  // Because buffer had space, ch <- 42 didn’t wait.

  **/
}

/*

len(ch) → current items in buffer
cap(ch) → max capacity

*/

/**

| Unbuffered  | Buffered       |
| ----------- | -------------- |
| Strong sync | Weak sync      |
| Handshake   | Queue          |
| Safer       | Faster         |
| Ordering    | Burst handling |

**/