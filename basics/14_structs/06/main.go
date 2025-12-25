package main

import "fmt"

func main(){
	// This is called an anonymous struct literal.
	language := struct{
		name string
		isGood bool
	}{"goalng", true}

	fmt.Println(language)
	/***
	  When do we use anonymous structs?
		✔ quick data
		✔ one-time use
		✔ tests / configs
		✔ small programs
	***/
}
