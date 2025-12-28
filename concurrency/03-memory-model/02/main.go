package main

import "fmt"

var count = 0

// 7️⃣ Data Race Problem (VERY IMPORTANT)
func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			count++
		}()
	}

	fmt.Println("count is", count)
	// 961
	// 969
	// 979
	/*
	What happens?
	  - Multiple goroutines modify count
	  - Leads to race condition
	*/

	// detect race using
	//go run -race main.go

}


/// Main Problem
/*
* 1. Data race → multiple goroutines modifying count
* 2. Main exits early → goroutines may not finish
*/

// why race happen
/*
count++
This looks atomic, but it is NOT.

It expands to:
  1. Read count
  2. Add 1
  3. Write back to count
 Multiple goroutines interleave these steps → lost updates.
 
*/