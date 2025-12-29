package main

import "fmt"

func main(){
  ch := make(chan int)

  go func ()  {
	ch <- 10 // channel need goroutine
  }()


  x := <-ch
  fmt.Println(x)

}

/***
---------- BLOCKING BEHAVIOURS -----------
1. This will deadlock
ch := make(chan int)
ch <- 1  // nobody recieving

2. Also deadlock
ch := make(chan int)
fmt.Println(<-ch) // no one sending

Send blocks until receive happens
Recieve blocks until send happens

// channel type
chan int
chan string
chan struct{}

Blocking = synchronization
(channels are powerful)

***/