package main

import "fmt"

const yo = "can decalre outside function! "


func main(){
	const name = "aditya kumar saini" // this can not be re-assigned
	const age = 22

	fmt.Println(yo, " ", age)

	/**************** constant grouping ********************/
	const (
		port = 5000
		host = "localhost"
	)

	// port = 6000 // it will complain as reassignment not allowed 

	fmt.Println(port, host)

}

