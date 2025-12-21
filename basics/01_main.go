package main

import "fmt"

func main(){
	/*********** simple values *****************************/

	// int
	fmt.Println(1)
	fmt.Println(1+2)

	// string
	fmt.Println("vande matram!")

	// boolean
	fmt.Println(true)
	fmt.Println(false)

	// float
	fmt.Println(10.5)
	fmt.Println(7.0 / 3.0)

	/*********************** Variables *********************/
	var name string = "raghu"
	var name2 = "aditya"   // infer the type

	var is_adult = true  // infer the type

	var age int = 89

	// short-hand syntax
	name3 := "golang"  // this synatx is used when decalre and assign value at same time
	var name4 string    // here when we only declare only (not assign) type defination is compulsary
	name4 = "hee hee"

	var price float32 = 50.3

	fmt.Println(name, "  " ,name2, "  ", is_adult, "  ", age, "  ", name3, "  ", name4, "  ", price)

}
