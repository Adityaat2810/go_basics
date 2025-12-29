package main
/*
Buffered Channel
-> buffered channel is a go channel that has capacity > 0

ch := make(chan int, 3) // buffered channel with capacity 3
- it can store values internally(like a queue)
- without requiring a receiver immediately

buffer vs unbuffer
buffer -> send block only when buffer is full
unbuffered -> send block until someone receives

Capacity = 3

[ 10 | 20 | 30 ]  ← buffer full


*/

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3  // deadlock( no reciever )

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// ✅ No goroutines needed

}

/**

--------- Blocking Behaviour -------------
1. send blocks when buffer is full
func main(){
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  ch <- 3 ( deadlock no reciever )

  
2. recive block when buffer is empty
func main(){
  ch := make(chan int, 2)
  fmt.Println(<-ch) (deadlock)
}

}
**/