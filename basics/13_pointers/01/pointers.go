package main

import "fmt"

/*
* A pointer is a variable that stores the memory address of another variable.
* Instead of holding a value directly, it points to where the value lives in memory.
 */

func main(){
//    x := 10
//    p := &x

   var x int = 10
   var p *int = &x  // * int type of pointer
   fmt.Println("memory adress of x is ", p, " and it holds value of ", x)

   /*
   * Dereferencing (* operator)
   * Dereferencing means accessing or modifying the value stored at the address.
   */

   *p = 20 // it cahnges value store at meory address p 
   fmt.Println("memory adress of x is ", p, " and it holds value of ", x)

}
