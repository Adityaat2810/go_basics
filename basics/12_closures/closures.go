package main

import "fmt"

// closures are functions which
// 1. are anonymous ( without name )
// 2. can access variables from outside their scope

func counter() func() int {
  var count int = 0

  return func() int {
	count += 1
	return  count
  }
}


func main(){
	increment := counter()
	for i:=0; i<10; i++ {
		fmt.Println(increment()) // due to closures it has acces to variable count of outer scope 
	}

}
