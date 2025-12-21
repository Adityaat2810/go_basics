package main

import (
	"fmt"
	"time"
)

func main(){
	i := 3

	// simple switch
	switch i {
	  case 1:
		fmt.Println("one")
	  case 2:
		fmt.Println("two")
	  case 3:
		fmt.Println("three")
	  case 4:
		fmt.Println("four")
	  case 5:
		fmt.Println("five")
	  default:
		fmt.Println("default")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	  case time.Saturday, time.Sunday:
		fmt.Println("It's weekend! ")
	  case time.Friday:
		fmt.Println("aaj friday hai!")
	  default:
		fmt.Println("today is ", time.Now().Weekday())
	}

	// type switch
	               // any type
	whoAmI := func (i interface{})  {
	  switch t := i.(type){
	    case int:
		   fmt.Println("It is a integer")
		case string:
			fmt.Println("It is a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("default", t)
	  }
	}

	whoAmI("aditya")
	whoAmI(21)
	whoAmI(true)
	whoAmI(78.7)
}
