package main

import "fmt"

/*
* A struct is a custom data type that groups related fields together.
* one real world entity = one struct
 */

type User struct {
    id int
    name string
    email string
    age int
}


func main(){
    var u User
    fmt.Println(u) // zero-value (struct 0, "", "", 0)

    // Field-wise initialization (recommended) {Order-independent}
    u2 := User{
	   id: 1,
	   name: "aditya",
	   email: "adityasainics10@gmail.com",
	   age: 21,
    }

	// Positional initialization
	u3 := User{2, "aditya", "adityasainica10@gamil.com", 12}


	u2.age = 1222  // accessing or updating


	fmt.Println(u2, u3.email)
}