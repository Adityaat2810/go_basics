package main

import "fmt"



func main(){
  /*
   all go-routines share same
    - same heap
	- share global variables
   This lead to data race condition
  */

  var x int

  go func() {
	x = 10
  }()

  for i:=0; i<500000; i=i+1{}

  fmt.Println(x)
}

/*
* ----- Race condition and go philosopy ----------
*
* Go philosophy:
* 	- "Do not communicate by sharing memory; share memory by communicating"
*
* Meaning
*  -> prefer channel
*  -> avaoid shared mutable state
*  -> use lock only when necessary
*/