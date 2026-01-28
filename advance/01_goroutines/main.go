package main

import (
	"fmt"
	"time"
)

/**
* Goroutines are just functions that leave the main thread and run in the background and come back to join the
* main thread once functions are finished/ready to return any value
*
* Goroutine do not stop program flow and are non blocking
*/

func main(){

  var err error
  fmt.Println("Beginning of the program!.")
  go sayHello()
  fmt.Println("After sayhello function.")

  // err = go doWork() // go routine do not return value like it

  go func ()  {
    err = doWork()
  }()

  go printNumbers()
  go printAlphabets()
  // these function print the numbers and alphabet in a interleaved
  // fashion-->

  // [Reason] The Go scheduler time-slices and switches between
  //  runnable goroutines whenever they block (sleep, I/O) or get preempted.





  time.Sleep(2* time.Second) // wait for goroutine to finish

  if err != nil{
    fmt.Println("Error: ", err)
  }else{
    fmt.Println("Work completed successfully")
  }

}

func sayHello(){
  time.Sleep(time.Second * 1) // sleep for 1 sec
  fmt.Println("hello from Goroutine!")

}

func printNumbers(){
  for i:=1; i<=5; i++{
    fmt.Println(i)
    time.Sleep(100 * time.Millisecond)
  }
}

func printAlphabets(){
  for ch:='A'; ch<='E'; ch++{
    fmt.Println(string(ch))
    time.Sleep(200 * time.Millisecond)
  }
}

func doWork() error{
  // simulate work
  time.Sleep(1* time.Second)

  return  fmt.Errorf("an error ocuured in do work")
}