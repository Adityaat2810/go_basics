package main

import "fmt"

// for -> only construct in go for looping
func main() {

	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop

	// for {
	// 	fmt.Println("1")
	// }

	// for loop
	fmt.Println("print 0 to 9")
	for j := 0; j<10; j++ {

		if ( j== 2){
			continue; // continue keyword
		}
		fmt.Println(j)
		if(j == 8) {
			break  // break keyword
		}
	}

	// range
	fmt.Println("print in range [0, 3)")

	for i := range(3) {
		fmt.Println(i)
	}
 
}