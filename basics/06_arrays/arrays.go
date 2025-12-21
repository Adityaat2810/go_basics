package main

import "fmt"

// array is numbered sequence of specific length

func main(){
  var nums [4]int  // declaration of integer array of length 4{ all filled with 0}

  // array length
  fmt.Println(len(nums))

  nums[0] = 11
  nums[1] = 12
  nums[2] = 13
  nums[3] = 14

  // accessing array elemnt
  for i:=0; i< len(nums); i++ {
	println("index ", i, " is ", nums[i])
  }

  // printing array
  fmt.Println("array is :-",nums)

  var vals[4]bool // by deafult filled with false
  var names[3]string // by default empty strings

  fmt.Println(vals, names)

  ///////// declartion & filling array at once //////////////////
  newNums := [3] int {1, 2, 3}
  fmt.Println(newNums)

  ////////// 2d arrays /////////////////////////
  nums2 := [2][2]int {{3, 4}, {5, 6}}
  fmt.Println(nums2)

}


/**
* On declartion auto filled values in array
* int->0
* string-> empty strings
* bool-> false

-> fixed size, that is predictable
-> memory optimization
-> constant time access
**/