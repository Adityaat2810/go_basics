package main

import "fmt"

// Unbuffered Channels
/*
BEHAVIOUR
 - Send and recieve must happen together
 - Acts like a handsake
*/

func main(){
  ch := make(chan int)

  go func ()  {
	fmt.Println("Sending!")
	ch<- 42
	fmt.Println("sent")
  }()

  fmt.Println("RECEIEVING")
  fmt.Println(<-ch)
  fmt.Println("RECIEVED")
  // Output order is guaranteed:

  /**
    Receiving
	Sending
	Sent
	Received

  **/

  // 👉 Used when ordering matters
}