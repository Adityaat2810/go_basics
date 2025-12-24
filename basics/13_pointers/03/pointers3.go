package main

import "fmt"

// Pointer to Pointer (**T)

func main(){
  x := 10
  p := &x
  pp := &p

  fmt.Println(p, pp)
  fmt.Println(**pp) // 10

}
