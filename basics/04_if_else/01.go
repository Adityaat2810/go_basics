package main

import "fmt"

func main() {
	age := 2

	// if-else
	if age < 18 {
		fmt.Println("person is not an adult")
	} else {
		fmt.Println("person is an adult")

	}

	// else if
	if age >= 18 {
		fmt.Println("person is adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is child")
	}

	// logical opertor
	var role = "admin"
	var hasPermissions = false

	if role == "admin" || hasPermissions { // conditional or
		fmt.Println("it will execute :)")
	} else {
		fmt.Println(" :(")
	}

	if age := 15; age > 18 {  // declaring varible within conditional
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("is teenager!", age)
	} else {
		fmt.Println("person is child")
	}

	fmt.Println("age is ", age)  // won't access outside conditional construct

	// go does not have ternary operator
	
}
