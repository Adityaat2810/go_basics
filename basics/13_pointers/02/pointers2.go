package main

import (
	"fmt"
)

// Why Pointers Exist (VERY IMPORTANT)

// 1. Without Pointer (Pass by Value)
func update(x int) {
    x = 100
}

func update2(x *int) {
    *x = 100
}

func main() {
    a := 10
	p := &a

    update(a)
    fmt.Println(a) // 10 { will return 10 as in function only value if passed}

	update2(p)
	fmt.Println(a) // 100 value is changed passed by refrence
	/*
	 *🔥 Use pointers when:
     1. You want to modify original data
     2. Avoid copying large structs
     3. Share state across functions
	*/

	/************** Zero Value of Pointer (nil) **************************/
	var p2 *int
    fmt.Println(p2) // nil
	// *p2 = 10 // panic: nil pointer dereference
    // fmt.Println(p2)

	if p2 != nil {  // Always check:
		fmt.Println("hey ",*p2)
	}


}

