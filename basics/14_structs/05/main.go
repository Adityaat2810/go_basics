package main

import "fmt"

type Address struct{
  city string
  pin int
}

type User struct{
  name string
  address Address
}

func main(){

	user := User{
		name: "aditya",
		address: Address{
			city: "meerut",
			pin: 250222,
		},
	}

	fmt.Println(user)

}