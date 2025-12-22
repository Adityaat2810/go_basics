package main

import "fmt"

// slices are dynamic array
// no need to give length
// most used construct in go
// + useful methods

func main(){
	///////////// declaration of slice ////////////////////////////
	// unitialized slice is nil (null)
	var nums[]int
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums)) // 0

	/////////////// slice with initial size ///////////////////////
	var nums2 = make([] int, 2) // it's dynamic with initial size = 2
	fmt.Println(nums2)
	fmt.Println(nums2 == nil) // false

	var nums3 = make([]int, 2, 5) // third parameter is for initial capacity
	fmt.Println(cap(nums3)) // capacity is maxm no. that can fit

	nums3 = append(nums3, 1)
	fmt.Println("slice :-", nums3, "capacity :-", cap(nums3)) // cap 5
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	fmt.Println("slice :-", nums3, "capacity :-", cap(nums3)) // cap 10 (resized to 10)
    nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 1)
	fmt.Println("slice :-", nums3, "capacity :-", cap(nums3)) // cap 20
	// when it exced the capapity it's get resized to 2*   //

	// var nums4 = make([]int , 0) // initial size 0 to esnure no 0, 0 elemnts initialized
	nums4 := []int {}
	nums4 = append(nums4, 10)
	println(len(nums4), nums4)





}
